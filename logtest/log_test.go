package logtest

import (
	"testing"
	"time"

	"github.com/func25/batchlog"
)

func TestMe(t *testing.T) {
	logger := batchlog.NewLogger()
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
