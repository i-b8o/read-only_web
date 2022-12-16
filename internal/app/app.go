package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	v1 "read-only_web/internal/controllers/http/v1"
	"read-only_web/internal/domain/service"
	"read-only_web/pkg/client/postgresql"

	"time"

	chapter_provider "read-only_web/internal/data_providers/db/postgresql/chapter"
	paragraph_provider "read-only_web/internal/data_providers/db/postgresql/paragraph"
	regulation_provider "read-only_web/internal/data_providers/db/postgresql/regulation"

	usecase_chapter "read-only_web/internal/domain/usecase/chapter"
	usecase_regulation "read-only_web/internal/domain/usecase/regulation"

	"read-only_web/internal/config"
	templateManager "read-only_web/internal/templmanager"

	"github.com/i-b8o/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	cfg        *config.Config
	router     *httprouter.Router
	logger     logging.Logger
	httpServer *http.Server
}

func NewApp(ctx context.Context, config *config.Config) (App, error) {
	logger := logging.GetLogger(config.AppConfig.LogLevel)

	logger.Print("router initializing")
	router := httprouter.New()
	logger.Print("swagger docs initializing")
	// hosting swagger specification
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	curdir, _ := os.Getwd()
	router.ServeFiles("/static/*filepath", http.Dir(curdir+"/internal/static/"))

	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)

	logger.Print("Postgres initializing")
	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("loading templates")
	templateManager := templateManager.NewTemplateManager(config.Template.Path)
	templateManager.LoadTemplates(ctx, logger)
	if err != nil {
		logger.Fatal(err)
	}

	regulationProvider := regulation_provider.NewRegulationStorage(pgClient)
	chapterProvider := chapter_provider.NewChapterStorage(pgClient)
	paragraphProvider := paragraph_provider.NewParagraphStorage(pgClient)

	regulationService := service.NewRegulationService(regulationProvider)
	chapterService := service.NewChapterService(chapterProvider)
	paragraphService := service.NewParagraphService(paragraphProvider)

	chapterUsecase := usecase_chapter.NewChapterUsecase(chapterService, paragraphService, regulationService, logger)
	regulationUsecase := usecase_regulation.NewRegulationUsecase(regulationService, chapterService, logger)

	chapterHandler := v1.NewChapterHandler(chapterUsecase, templateManager)
	regulationHandler := v1.NewRegulationHandler(regulationUsecase, templateManager)

	regulationHandler.Register(router)
	chapterHandler.Register(router)

	return App{cfg: config, router: router, logger: logger}, nil
}

func (a *App) Run(ctx context.Context) error {

	// Define the listener (Unix or TCP)
	// var listener net.Listener

	a.logger.Infof("bind application to host: %s and port: %s", a.cfg.HTTP.IP, a.cfg.HTTP.Port)
	var err error
	// start up a tcp listener
	// listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		a.logger.Fatal(err)
	}

	// create a new Cors handler
	c := cors.New(cors.Options{
		AllowedMethods: a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins: a.cfg.HTTP.CORS.AllowedOrigins,
		AllowedHeaders: a.cfg.HTTP.CORS.AllowedHeaders,
	})

	// apply the CORS specification on the request, and add relevant CORS headers
	handler := c.Handler(a.router)

	// define parameters for an HTTP server
	a.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", a.cfg.HTTP.Port),
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	a.logger.Println("application initialized and started")

	// accept incoming connections on the listener, creating a new service goroutine for each
	if err := a.httpServer.ListenAndServeTLS("/etc/ssl/certs/read-only.crt", "/etc/ssl/certs/read-only.key"); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			a.logger.Warn("server shutdown")

		default:
			a.logger.Fatal(err)
		}
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		a.logger.Fatal(err)
	}
	return nil
}
