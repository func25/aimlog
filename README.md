# batchlog

batchlog is used to batch the logs (built on top of zerolog), avoid spamming 

* [Installation](#installation)
* [Samples](#samples)
 	* [Batching](#batching)
	* [Grouping](#grouping)
* [Status](#status-pre-release)


## Installation

`go get github.com/func25/batchlog`

## Samples

### Batching

Batch with three variables: "tokenId", "isBatch" and "message"
```go
logger := batchlog.NewLogger()
logger.Debug().BatchStr("tokenId", "1").BatchBool("isBatch", false).BatchMsg("hello")
logger.Debug().BatchStr("tokenId", "1").BatchMsg("hello")
logger.Debug().BatchStr("tokenId", "2").BatchMsg("hello")
logger.Debug().BatchStr("tokenId", "2").BatchMsg("hello")
logger.Debug().BatchStr("tokenId", "2").BatchMsg("hello")
logger.Debug().BatchStr("tokenId", "2").BatchMsg("hellok")
logger.Debug().BatchStr("tokenId", "2").BatchMsg("hello")
logger.Error().BatchStr("tokenId", "2").BatchMsg("hello")
logger.Debug().BatchStr("tokenId", "2").BatchMsg("hello")
logger.Debug().BatchStr("tokenId", "2").BatchMsg("hello")
logger.Info().BatchStr("tokenId", "2").BatchMsg("hello")
```

The result will be:
```
{"level":"info","tokenId":"2","message":"hello","__repeat":1}
{"level":"error","tokenId":"2","message":"hello","__repeat":1}
{"level":"debug","tokenId":"2","message":"hellok","__repeat":1}
{"level":"debug","tokenId":"2","message":"hello","__repeat":6}
{"level":"debug","tokenId":"1","message":"hello","__repeat":1}
{"level":"debug","tokenId":"1","isBatch":false,"message":"hello","__repeat":1}
```

### Grouping

This sample will batch first 20 messages in 1 batch and also create group of ids for the log
```go
logger := batchlog.NewLogger(batchlog.OptTimeout(time.Hour))
for i := 0; i < 30; i++ {
	time.Sleep(1500 * time.Millisecond)
	logger.Debug().BatchStr("tokenID", "123456").GroupInt("id", i).BatchMsg("hello")
}
```

```go
{"level":"debug","tokenID":"123456","message":"hello","__repeat":20,"id":["0","1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19"]}
{"level":"debug","tokenID":"123456","message":"hello","__repeat":10,"id":["20","21","22","23","24","25","26","27","28","29"]}
```

## Status: pre-release
This lib is under developing, please notice when using it
