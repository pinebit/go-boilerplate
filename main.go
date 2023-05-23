package main

import (
	"context"

	"github.com/pinebit/go-boilerplate/config"
	"github.com/pinebit/go-boilerplate/logger"
	"github.com/pinebit/go-boilerplate/services/http"
	"go.uber.org/dig"

	"github.com/pinebit/go-boot/boot"
)

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(func() *config.Config {
		config := config.NewDefaultConfig()
		if err := config.LoadFromToml("config/config.toml"); err != nil {
			panic(err)
		}
		return config
	})
	container.Provide(func(config *config.Config) logger.Logger {
		logger, err := logger.NewLogger(config.DevMode)
		if err != nil {
			panic(err)
		}
		return logger
	})
	container.Provide(http.NewRouter)
	container.Provide(http.NewServer)
	container.Provide(func(server http.Server, config *config.Config) boot.Applicaiton {
		return boot.NewApplicationForService(server, config.ShutdownTimeout.Duration())
	})

	return container
}

func main() {
	container := buildContainer()

	err := container.Invoke(func(app boot.Applicaiton, logger logger.Logger) {
		logger.Info("Starting server application...")

		if err := app.Run(context.Background()); err != nil {
			logger.Errorw("Server application shutdown error", "err", err)
		} else {
			logger.Info("Server application stopped gracefully.")
		}
	})

	if err != nil {
		panic(err)
	}
}
