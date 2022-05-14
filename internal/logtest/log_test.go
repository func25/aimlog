package logtest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/func25/batchlog"
)

func TestMe(t *testing.T) {
	logger := batchlog.New(os.Stdout)
	logger.Debug().BatchStr("tokenID", "123456").BatchBool("isBatch", false).BatchMsg("hello")
	logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
	logger.Debug().BatchStr("tokenID", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenID", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenID", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenID", "123457").BatchMsg("hellok")
	logger.Debug().BatchStr("tokenID", "123457").BatchMsg("hello")
	logger.Error().BatchStr("tokenID", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenID", "123457").BatchMsg("hello")
	logger.Debug().BatchStr("tokenID", "123457").BatchMsg("hello")
	logger.Info().BatchStr("tokenID", "123457").BatchMsg("hello")
	time.Sleep(1000 * time.Second)
}

func TestWait(t *testing.T) {
	logger := batchlog.New(os.Stdout, batchlog.OptTimeout(1*time.Hour))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("log")
		logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
	}
	fmt.Println("done")
	time.Sleep(10 * time.Hour)
}

func TestTimeout(t *testing.T) {
	logger := batchlog.New(os.Stdout, batchlog.OptTimeout(5*time.Second))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("log")
		logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
	}
	fmt.Println("done")
	time.Sleep(10 * time.Hour)
}

func TestMaxRelatitveBatch(t *testing.T) {
	logger := batchlog.New(os.Stdout, batchlog.OptTimeout(time.Hour))
	for i := 0; i < 30; i++ {
		time.Sleep(1500 * time.Millisecond)
		logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
	}
	fmt.Println("done")
	time.Sleep(time.Hour)
}

func TestGroup(t *testing.T) {
	logger := batchlog.New(os.Stdout, batchlog.OptTimeout(time.Hour))
	for i := 0; i < 30; i++ {
		time.Sleep(1500 * time.Millisecond)
		logger.Debug().BatchStr("tokenID", "123456").GroupInt("id", i).BatchMsg("hello")
	}
	fmt.Println("done")
	time.Sleep(time.Hour)
}

func TestScenario(t *testing.T) {
	logger := batchlog.New(os.Stdout, batchlog.OptTimeout(time.Hour))

	go func() {
		logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchMsg("hello")
		time.Sleep(time.Second)
		fmt.Println("done 1")
	}()

	go func() {
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchMsg("hello")
		time.Sleep(time.Second)
		fmt.Println("done 2")
	}()

	go func() {
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		logger.Debug().BatchStr("tokenID", "123456").BatchBool("bool", false).BatchInt("int", 1).BatchMsg("hello")
		time.Sleep(time.Second)
		fmt.Println("done 3")
	}()

	for {
		time.Sleep(time.Second)
	}
}

func TestContext(t *testing.T) {
	logger := batchlog.FromZLog(batchlog.NewZLog(os.Stdout).With().Str("op", "FuseHero").Logger())
	logger.Info().BatchAny("test", "this is test").Send()
	logger.Info().BatchAny("test", "this is test").Send()
	logger.Info().BatchAny("test", "this is test").Send()
	time.Sleep(15 * time.Second)
}
