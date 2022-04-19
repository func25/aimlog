package logtest

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	l zerolog.Logger

	Timeout  time.Duration
	MaxBatch int
}

func NewLogger() Logger {
	return Logger{
		l:        zerolog.New(os.Stdout),
		Timeout:  2 * time.Second,
		MaxBatch: 10,
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
