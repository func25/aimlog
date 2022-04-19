package batchlog

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	l zerolog.Logger
}

func NewLogger() Logger {
	return Logger{
		l: zerolog.New(os.Stdout),
	}
}

func (l *Logger) Debug() *event {
	return newRawEvent(l.l.Debug, zerolog.DebugLevel)
}

func (l *Logger) Info() *event {
	return newRawEvent(l.l.Info, zerolog.InfoLevel)
}

func (l *Logger) Warn() *event {
	return newRawEvent(l.l.Warn, zerolog.WarnLevel)
}

func (l *Logger) Error() *event {
	return newRawEvent(l.l.Error, zerolog.ErrorLevel)
}

func (l *Logger) Panic() *event {
	return newRawEvent(l.l.Panic, zerolog.PanicLevel)
}

func (l *Logger) Fatal() *event {
	return newRawEvent(l.l.Fatal, zerolog.FatalLevel)
}
