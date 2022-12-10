package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	v1 "read-only_web/internal/controllers/http/v1"
	"read-only_web/internal/domain/service"

	"time"

	chapter_adapter "read-only_web/internal/adapters/grpc/v1/chapter"
	paragraph_adapter "read-only_web/internal/adapters/grpc/v1/paragraph"
	regulation_adapter "read-only_web/internal/adapters/grpc/v1/regulation"
	usecase_chapter "read-only_web/internal/domain/usecase/chapter"
	usecase_regulation "read-only_web/internal/domain/usecase/regulation"

	"read-only_web/internal/config"
	templateManager "read-only_web/internal/templmanager"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
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

	logger.Print("heartbeat initializing")

	logger.Print("Postgres initializing")
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", config.Reader.IP, config.Reader.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return App{}, err
	}
	regulationGrpcClient := pb.NewRegulationGRPCClient(conn)
	chapterGrpcClient := pb.NewChapterGRPCClient(conn)
	paragraphGrpcClient := pb.NewParagraphGRPCClient(conn)

	logger.Print("loading templates")
	templateManager := templateManager.NewTemplateManager(config.Template.Path)
	templateManager.LoadTemplates(ctx, logger)
	if err != nil {
		logger.Fatal(err)
	}

	regulationAdapter := regulation_adapter.NewRegulationStorage(regulationGrpcClient)
	chapterAdapter := chapter_adapter.NewChapterStorage(chapterGrpcClient)
	paragraphAdapter := paragraph_adapter.NewChapterStorage(paragraphGrpcClient)

	regulationService := service.NewRegulationService(regulationAdapter)
	chapterService := service.NewChapterService(chapterAdapter)
	paragraphService := service.NewParagraphService(paragraphAdapter)

	// paragraphUsecase := paragraph_usecase.NewParagraphUsecase(paragraphService, chapterService, linkService, speechService)
	chapterUsecase := usecase_chapter.NewChapterUsecase(chapterService, paragraphService, regulationService, logger)
	regulationUsecase := usecase_regulation.NewRegulationUsecase(regulationService, chapterService, logger)
	// searchUsecase := search_usecase.NewSearchUsecase(searchService)

	// paragraphHandler := v1.NewParagraphHandler(paragraphUsecase, config.HTTP.UseToInsertData)
	chapterHandler := v1.NewChapterHandler(chapterUsecase, templateManager)
	regulationHandler := v1.NewRegulationHandler(regulationUsecase, templateManager)
	// searchHandler := v1.NewSearchHandler(searchUsecase)

	regulationHandler.Register(router)
	chapterHandler.Register(router)

	// read ca's cert, verify to client's certificate
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// caPem, err := ioutil.ReadFile(homeDir + "/certs/ca-cert.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // create cert pool and append ca's cert
	// certPool := x509.NewCertPool()
	// if !certPool.AppendCertsFromPEM(caPem) {
	// 	log.Fatal(err)
	// }

	// // read server cert & key
	// serverCert, err := tls.LoadX509KeyPair(homeDir+"/certs/server-cert.pem", homeDir+"/certs/server-key.pem")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // configuration of the certificate what we want to
	// conf := &tls.Config{
	// 	Certificates: []tls.Certificate{serverCert},
	// 	ClientAuth:   tls.RequireAndVerifyClientCert,
	// 	ClientCAs:    certPool,
	// }

	// //create tls certificate
	// tlsCredentials := credentials.NewTLS(conf)

	// grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	// grpcServer := grpc.NewServer()
	// server := grpc_service.NewRegulationGRPCService(regulationUsecase, chapterUsecase, paragraphUsecase)
	// pb.RegisterRegulationGRPCServer(grpcServer, server)

	return App{cfg: config, router: router, logger: logger}, nil
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP(ctx)
	})
	return grp.Wait()
}

func (a *App) startHTTP(ctx context.Context) error {

	// Define the listener (Unix or TCP)
	var listener net.Listener

	a.logger.Infof("bind application to host: %s and port: %s", a.cfg.HTTP.IP, a.cfg.HTTP.Port)
	var err error
	// start up a tcp listener
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		a.logger.Fatal(err)
	}

	// create a new Cors handler
	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost},
		AllowedOrigins:     []string{"http://localhost:10000"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Access-Token", "Refresh-Token", "Location", "Authorization", "Content-Disposition"},
		Debug:              false,
	})

	// apply the CORS specification on the request, and add relevant CORS headers
	handler := c.Handler(a.router)

	// define parameters for an HTTP server
	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	a.logger.Println("application initialized and started")

	// accept incoming connections on the listener, creating a new service goroutine for each
	if err := a.httpServer.Serve(listener); err != nil {
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
