package batchlog

import (
	"io"

	"github.com/rs/zerolog"
)

func NewZLog(w io.Writer) zerolog.Logger {
	return zerolog.New(w)
}

func FromZLog(logger zerolog.Logger, opts ...BatchOption) Logger {
	return Logger{
		logger: logger,
		opts:   opts,
	}
}
