package logger

import (
	"io"
	"log/slog"
	"os"
)

// NewLogger uses readable debug logs in development and JSON logs in production.
func NewLogger(devMode bool) *slog.Logger {
	return newLogger(os.Stdout, devMode)
}

func newLogger(output io.Writer, devMode bool) *slog.Logger {
	if devMode {
		return slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return slog.New(slog.NewJSONHandler(output, nil))
}
