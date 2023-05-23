package main

import (
	"context"
	net_http "net/http"

	"github.com/pinebit/go-boilerplate/config"
	"github.com/pinebit/go-boilerplate/logger"
	"github.com/pinebit/go-boilerplate/services/http"

	"github.com/pinebit/go-boot/boot"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	net_http.Handle("/metrics", promhttp.Handler())

	httpServer := http.NewServer(config.HttpServer)

	logger.Info("Starting server application...")

	application := boot.NewApplicationForService(httpServer, config.ShutdownTimeout.Duration())
	if err = application.Run(context.Background()); err != nil {
		logger.Fatalw("Server application shutdown error", "err", err)
	}

	logger.Info("Server application stopped gracefully.")
}
