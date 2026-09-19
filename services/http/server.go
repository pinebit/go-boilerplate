package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/pinebit/go-boilerplate/config"
)

type Server struct {
	server *http.Server
	log    *slog.Logger
}

func NewServer(log *slog.Logger, cfg *config.Config, handler http.Handler) *Server {
	return &Server{log: log, server: &http.Server{
		Addr:              net.JoinHostPort(cfg.HttpServer.Address, strconv.Itoa(int(cfg.HttpServer.Port))),
		Handler:           handler,
		ReadHeaderTimeout: cfg.HttpServer.ReadHeaderTimeout.Duration(),
		ReadTimeout:       cfg.HttpServer.ReadTimeout.Duration(),
		WriteTimeout:      cfg.HttpServer.WriteTimeout.Duration(),
		IdleTimeout:       cfg.HttpServer.IdleTimeout.Duration(),
		ErrorLog:          slog.NewLogLogger(log.Handler(), slog.LevelError),
	}}
}

// Run stops accepting requests on cancellation and allows active requests to drain.
func (s *Server) Run(ctx context.Context, shutdownTimeout time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	listener, err := (&net.ListenConfig{}).Listen(ctx, "tcp", s.server.Addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	s.log.Info("HTTP server started", "address", listener.Addr().String())
	return s.serve(ctx, listener, shutdownTimeout)
}

func (s *Server) serve(ctx context.Context, listener net.Listener, shutdownTimeout time.Duration) error {
	done := make(chan error, 1)
	go func() { done <- s.server.Serve(listener) }()
	select {
	case err := <-done:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("serve: %w", err)
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		err := s.server.Shutdown(shutdownCtx)
		if err != nil {
			err = errors.Join(err, s.server.Close())
		}
		serveErr := <-done
		if !errors.Is(serveErr, http.ErrServerClosed) {
			err = errors.Join(err, serveErr)
		}
		if err != nil {
			return fmt.Errorf("shutdown: %w", err)
		}
		s.log.Info("HTTP server stopped")
		return nil
	}
}
