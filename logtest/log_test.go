package logtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/func25/batchlog"
	"github.com/func25/gofunc/mathfunc"
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

func TestWait(t *testing.T) {
	logger := batchlog.NewLogger(batchlog.OptTimeout(1 * time.Hour))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("log")
		logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
	}
	fmt.Println("done")
	time.Sleep(10 * time.Hour)
}

func TestTimeout(t *testing.T) {
	logger := batchlog.NewLogger(batchlog.OptTimeout(5 * time.Second))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("log")
		logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
	}
	fmt.Println("done")
	time.Sleep(10 * time.Hour)
}

func TestMaxRelatitveBatch(t *testing.T) {
	logger := batchlog.NewLogger(batchlog.OptTimeout(time.Hour))
	for i := 0; i < 30; i++ {
		time.Sleep(1500 * time.Millisecond)
		logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
	}
	fmt.Println("done")
	time.Sleep(time.Hour)
}

func TestScenario(t *testing.T) {
	logger := batchlog.NewLogger(batchlog.OptTimeout(time.Hour))

	go func() {
		logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchMsg("hello")
		batchlog.Test()
		fmt.Println("done 1")
	}()

	go func() {
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchMsg("hello")
		batchlog.Test()
		fmt.Println("done 2")
	}()

	go func() {
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		logger.Debug().BatchStr("tokenId", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		batchlog.Test()
		fmt.Println("done 3")
	}()

	for {
		batchlog.Test()
	}
}
