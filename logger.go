package batchlog

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
	opts   []BatchOption
}

func New(w io.Writer, opts ...BatchOption) Logger {
	return Logger{
		logger: zerolog.New(w),
		opts:   opts,
	}
}

func (l *Logger) Debug() *event {
	return newRawEvent(l, l.logger.Debug, zerolog.DebugLevel)
}

func (l *Logger) Info() *event {
	return newRawEvent(l, l.logger.Info, zerolog.InfoLevel)
}

func (l *Logger) Warn() *event {
	return newRawEvent(l, l.logger.Warn, zerolog.WarnLevel)
}

func (l *Logger) Error() *event {
	return newRawEvent(l, l.logger.Error, zerolog.ErrorLevel)
}

func (l *Logger) Panic() *event {
	return newRawEvent(l, l.logger.Panic, zerolog.PanicLevel)
}

func (l *Logger) Fatal() *event {
	return newRawEvent(l, l.logger.Fatal, zerolog.FatalLevel)
}
