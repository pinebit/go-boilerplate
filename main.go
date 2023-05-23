package main

import (
	"context"

	"github.com/pinebit/go-boilerplate/config"
	"github.com/pinebit/go-boilerplate/logger"
	"github.com/pinebit/go-boilerplate/services/http"

	"github.com/pinebit/go-boot/boot"
)

func main() {
	config := config.NewDefaultConfig()
	if err := config.LoadFromToml("config/config.toml"); err != nil {
		panic(err)
	}

	logger, err := logger.NewLogger(config.DevMode)
	if err != nil {
		panic(err)
	}

	router := http.NewRouter(logger, config)
	httpServer := http.NewServer(logger, config.HttpServer, router)

	logger.Info("Starting server application...")

	application := boot.NewApplicationForService(httpServer, config.ShutdownTimeout.Duration())
	if err = application.Run(context.Background()); err != nil {
		logger.Fatalw("Server application shutdown error", "err", err)
	}

	logger.Info("Server application stopped gracefully.")
}
