package batchlog

import (
	"encoding/json"
	"strconv"
)

func (e *event) GroupBool(key string, value bool) *event {
	e.group = pairString{
		key:   key,
		value: strconv.FormatBool(value),
	}
	return e
}

func (e *event) GroupInt(key string, value int) *event {
	e.group = pairString{
		key:   key,
		value: strconv.FormatInt(int64(value), 10),
	}
	return e
}

func (e *event) GroupFloat32(key string, value float32) *event {
	e.group = pairString{
		key:   key,
		value: strconv.FormatFloat(float64(value), 'f', -1, 32),
	}
	return e
}

func (e *event) GroupFloat64(key string, value float64) *event {
	e.group = pairString{
		key:   key,
		value: strconv.FormatFloat(float64(value), 'f', -1, 64),
	}
	return e
}

func (e *event) GroupInterface(key string, i interface{}) *event {
	if b, err := json.Marshal(i); err != nil {
		return nil
	} else {
		e.group = pairString{
			key:   key,
			value: string(b),
		}
		return e
	}
}

func (e *event) GroupStr(key string, value string) *event {
	e.group = pairString{
		key:   key,
		value: value,
	}
	return e
}

func (e *event) GroupErr(err error) *event {
	e.group = pairString{
		key:   "error",
		value: err.Error(),
	}
	return e
}
