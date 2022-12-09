package main

// TODO 404 page
// TODO SWAGGER

import (
	"context"

	"read-only_web/internal/app"
	"read-only_web/internal/config"

	"github.com/i-b8o/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.GetConfig()
	logger := logging.GetTelegramLogger(ctx, "reg_be", cfg.Logger.LoggerUsername, cfg.Logger.Token, cfg.Logger.ChatID)

	ctx = logging.ContextWithLogger(ctx, logger)

	logger.Info("config initializing")

	logger.Info("logger initializing")

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Running Application")

	a.Run(ctx)
}
