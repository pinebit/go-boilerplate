package http

import (
	"context"
	"fmt"
	"net"
	net_http "net/http"

	"github.com/pinebit/go-boilerplate/config"

	"github.com/pinebit/go-boot/boot"
)

type Server interface {
	boot.Service
}

type server struct {
	cancelBaseCtx context.CancelFunc
	serverService boot.HttpServer
}

func NewServer(config *config.HttpServerConfig) Server {
	ctx, cancel := context.WithCancel(context.Background())
	netHttpServer := &net_http.Server{
		Addr: fmt.Sprintf("%s:%d", config.Address, config.Port),
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}
	return &server{
		cancelBaseCtx: cancel,
		serverService: boot.NewHttpServer(netHttpServer),
	}
}

func (s server) Start(ctx context.Context) error {
	return s.serverService.Start(ctx)
}

func (s server) Stop(ctx context.Context) error {
	s.cancelBaseCtx()
	return s.serverService.Stop(ctx)
}
