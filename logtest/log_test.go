package logtest

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/rs/zerolog"
)

type loggerPool struct {
	pool map[string]int
}

type job struct {
}

var pool *loggerPool

func init() {
	pool = &loggerPool{
		make(map[string]int, 100),
	}
}

func TestMe(t *testing.T) {
	logger := NewLogger()
	fmt.Println(unsafe.Sizeof(logger))
	msg := logger.Debug().Str("H", "hello guy")
	msg.Msg("message")
}

func captainHook(e *zerolog.Event, level zerolog.Level, message string) {

}
