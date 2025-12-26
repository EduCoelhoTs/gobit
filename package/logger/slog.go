package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	Env   string
	Level slog.Level
}

func New(cfg Config) *slog.Logger {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level:     cfg.Level,
		AddSource: cfg.Env == "dev",
	}

	if cfg.Env == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
