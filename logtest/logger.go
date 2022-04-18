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
	return newRawEvent(l.l.Debug)
}

func (l *Logger) Info() *event {
	return newRawEvent(l.l.Info)
}

func (l *Logger) Warn() *event {
	return newRawEvent(l.l.Warn)
}

func (l *Logger) Error() *event {
	return newRawEvent(l.l.Error)
}

func (l *Logger) Panic() *event {
	return newRawEvent(l.l.Panic)
}

func (l *Logger) Fatal() *event {
	return newRawEvent(l.l.Fatal)
}
