package logtest

import (
	"testing"
	"time"
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
	logger.Debug().BatchStr("tokenId", "123456").BatchBool("isBatch", false).BatchMsg("hello")
	logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
	logger.Debug().BatchStr("tokenId", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenId", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenId", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenId", "123457").BatchMsg("hellok")
	logger.Debug().BatchStr("tokenId", "123457").BatchMsg("hello")
	logger.Error().BatchStr("tokenId", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenId", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenId", "123457").BatchMsg("hello")
	logger.Info().BatchStr("tokenId", "123457").BatchMsg("hello")
	time.Sleep(1000 * time.Second)
}
