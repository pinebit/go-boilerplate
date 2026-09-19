package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/pinebit/go-boilerplate/config"
	"github.com/pinebit/go-boilerplate/logger"
	httpservice "github.com/pinebit/go-boilerplate/services/http"
)

func run(ctx context.Context, configPath string) error {
	cfg := config.NewDefaultConfig()
	if err := cfg.LoadFromToml(configPath); err != nil {
		return err
	}
	log := logger.NewLogger(cfg.DevMode)
	server := httpservice.NewServer(log, cfg, httpservice.NewRouter(log))
	return server.Run(ctx, cfg.ShutdownTimeout.Duration())
}

func main() {
	configPath := flag.String("config", "config/config.toml", "path to TOML configuration")
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := run(ctx, *configPath); err != nil {
		slog.Error("Application failed", "error", err)
		os.Exit(1)
	}
}
