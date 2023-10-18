package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger() *Logger {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	return &Logger{
		logger: l,
	}
}

func (l *Logger) Error(err error, args ...any) {
	l.logger.Error(err.Error(), "data", args)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, "data", args)
}
