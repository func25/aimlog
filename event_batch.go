package batchlog

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rs/zerolog"
)

type pairString struct {
	key   string
	value string
}

// event is based on zerolog.Event, so dont reuse this after logging
type event struct {
	logger *Logger

	level zerolog.Level
	event *zerolog.Event
	done  bool

	batchKeysM map[string]bool // map of keys that need to batched
	batchKeysA []string        // array of keys that need to batched

	group pairString
}

func newRawEvent(l *Logger, e func() *zerolog.Event, lvl zerolog.Level) *event {
	return &event{
		logger: l,
		level:  lvl,
		event:  e(),
		// groupKey:   "",
		batchKeysM: make(map[string]bool),
		batchKeysA: make([]string, 0),
	}
}

// BatchMsg will batch the message (string) and logging out
// this equal: e.BatchStr("message", msg) + e.Send()
//
// NOTICE: once this method is called, the *event should be disposed.
// Calling Msg twice can have unexpected result.
func (e *event) BatchMsg(msg string) {
	if e.done {
		return
	}
	e.done = true

	if len(msg) > 0 {
		e.BatchStr("message", msg)
	}
	batcher.Batch(e)
}

// Msg sends the *event with msg added as the message field if not empty.
// it will processed as async batch operation if any param of it is batched
// otherwise it just prints out immediately
//
// NOTICE: once this method is called, the *event should be disposed.
// Calling Msg twice can have unexpected result.
func (e *event) Msg(msg string) {
	if e.done {
		return
	}
	e.done = true

	batcher.Batch(e.Str("message", msg))
}

// Send is equivalent to calling Msg("").
//
// NOTICE: once this method is called, the *Event should be disposed.
func (e *event) Send() {
	if e.done {
		return
	}
	e.done = true

	batcher.Batch(e)
}

func (e *event) BatchBool(key string, value bool) *event {
	if e.Bool(key, value) == nil {
		return nil
	}

	return e.batch(key, strconv.FormatBool(value))
}

func (e *event) BatchInt(key string, value int) *event {
	if e.Int(key, value) == nil {
		return nil
	}

	return e.batch(key, strconv.FormatInt(int64(value), 10))
}

func (e *event) BatchFloat32(key string, value float32) *event {
	if e.Float32(key, value) == nil {
		return nil
	}

	return e.batch(key, strconv.FormatFloat(float64(value), 'f', -1, 32))
}

func (e *event) BatchFloat64(key string, value float64) *event {
	if e.Float64(key, value) == nil {
		return nil
	}
	return e.batch(key, strconv.FormatFloat(float64(value), 'f', -1, 64))
}

func (e *event) BatchInterface(key string, i interface{}) *event {
	if e.Interface(key, i) == nil {
		return nil
	}

	if b, err := json.Marshal(i); err != nil {
		return nil
	} else {
		return e.batch(key, string(b))
	}
}

func (e *event) BatchStr(key string, value string) *event {
	if e.Str(key, value) == nil {
		return nil
	}

	return e.batch(key, value)
}

func (e *event) BatchErr(err error) *event {
	if e.Err(err) == nil {
		return nil
	}

	return e.batch("error", err.Error())
}

func (e *event) batch(key string, value string) *event {
	realKey := fmt.Sprintf("%s_%s", key, value)
	e.batchKeysM[realKey] = true
	e.batchKeysA = append(e.batchKeysA, realKey)
	return e
}
