package http

import (
	"context"
	"fmt"
	"net"
	net_http "net/http"

	"github.com/pinebit/go-boilerplate/config"
	"github.com/pinebit/go-boilerplate/logger"
	"github.com/pinebit/go-boot/boot"
)

type Server interface {
	boot.Service
}

type server struct {
	logger        logger.Logger
	cancelBaseCtx context.CancelFunc
	serverService boot.HttpServer
}

func NewServer(logger logger.Logger, config *config.HttpServerConfig, router Router) Server {
	ctx, cancel := context.WithCancel(context.Background())
	netHttpServer := &net_http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Address, config.Port),
		Handler: router.Handler(),
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}
	return &server{
		logger:        logger.Named("http.Server"),
		cancelBaseCtx: cancel,
		serverService: boot.NewHttpServer(netHttpServer),
	}
}

func (s server) Start(ctx context.Context) error {
	s.logger.Debug("Starting http server...")
	return s.serverService.Start(ctx)
}

func (s server) Stop(ctx context.Context) error {
	s.logger.Debug("Stoppping http server...")
	s.cancelBaseCtx()
	return s.serverService.Stop(ctx)
}
