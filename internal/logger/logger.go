package logger

import (
	"log/slog"
	"os"
	"strings"
)

func RegisterLogger(logLevel string) {

	// default log
	opt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	//
	// NOTE: Reminder Go doesn't have fall through case
	//       make sure all cases are filled out correctly
	//
	switch strings.ToLower(logLevel) {
	case "debug":
		opt.Level = slog.LevelDebug

	case "warn":
		opt.Level = slog.LevelWarn

	case "error":
		opt.Level = slog.LevelError

	case "info":
		opt.Level = slog.LevelInfo

	default:
		opt.Level = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	slog.SetDefault(logger)
}

func NewForcedLogger() *slog.Logger {
	opt := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	l := slog.New(slog.NewJSONHandler(os.Stdout, opt))
	return l
}
