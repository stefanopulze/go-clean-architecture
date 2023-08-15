package logging

import (
	"go-clean-architecture/internal/config"
	"golang.org/x/exp/slog"
	"os"
	"strings"
)

var logger *slog.Logger

type Logger struct {
	*slog.Logger
}

func Init(opts config.LogOpts) *Logger {
	var handler slog.Handler
	logConfig := &slog.HandlerOptions{
		Level: logLevel(opts.Level),
	}

	if opts.Type == "console" {
		handler = slog.NewTextHandler(os.Stdout, logConfig)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, logConfig)
	}

	logger = slog.New(handler)
	slog.SetDefault(logger)

	return &Logger{
		logger,
	}
}

func logLevel(l string) slog.Level {
	switch strings.ToLower(l) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
