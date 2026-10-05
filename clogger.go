// Package clogger is a small HTTP package for logging requests
package clogger

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	log *slog.Logger
}

func New(l *slog.Logger) *Logger {
	if l == nil {
		l = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return &Logger{log: l}
}

func logLevel(status int) slog.Level {
	switch {
	case status >= 500:
		return slog.LevelError
	case status >= 400:
		return slog.LevelWarn
	case status >= 300:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

func (log *Logger) LogRequest(method, url string, status int, duration time.Duration) {
	log.log.Log(context.Background(), logLevel(status), "request", "url", url, "method", method, "status", status, "duration", duration)
}
