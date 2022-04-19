package batchlog

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	l    zerolog.Logger
	opts []BatchOption
}

func NewLogger(opts ...BatchOption) Logger {
	return Logger{
		l:    zerolog.New(os.Stdout),
		opts: opts,
	}
}

func (l *Logger) Debug() *event {
	return newRawEvent(l, l.l.Debug, zerolog.DebugLevel)
}

func (l *Logger) Info() *event {
	return newRawEvent(l, l.l.Info, zerolog.InfoLevel)
}

func (l *Logger) Warn() *event {
	return newRawEvent(l, l.l.Warn, zerolog.WarnLevel)
}

func (l *Logger) Error() *event {
	return newRawEvent(l, l.l.Error, zerolog.ErrorLevel)
}

func (l *Logger) Panic() *event {
	return newRawEvent(l, l.l.Panic, zerolog.PanicLevel)
}

func (l *Logger) Fatal() *event {
	return newRawEvent(l, l.l.Fatal, zerolog.FatalLevel)
}
