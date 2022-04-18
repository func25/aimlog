package logtest

import (
	"encoding/json"
	"strconv"

	"github.com/rs/zerolog"
)

type event struct {
	*zerolog.Event
	repeat   int
	groupKey string

	batchKeysM map[string]bool
	batchKeysA []string
	checkMap   map[string]string
}

func newRawEvent(e func() *zerolog.Event) *event {
	return &event{
		Event:      e(),
		repeat:     0,
		groupKey:   "",
		batchKeysM: make(map[string]bool),
		batchKeysA: make([]string, 0),
	}
}

func (e *event) BatchMsg(msg string) {

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
	e.checkMap[key] = value
	e.batchKeysM[key] = true
	e.batchKeysA = append(e.batchKeysA, key)
	return e
}
