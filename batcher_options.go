package batchlog

import (
	"time"
)

type BatchOption func(*chainData)

func (c *chainData) applyOpts(opts ...BatchOption) *chainData {
	for i := range opts {
		opts[i](c)
	}

	return c
}

func OptTimeout(timeout time.Duration) BatchOption {
	return func(q *chainData) {
		q.timeout = timeout
	}
}

func OptMaxRelativeBatch(maxRelativeBatch int) BatchOption {
	return func(q *chainData) {
		q.maxRelativeBatch = maxRelativeBatch
	}
}

func OptWait(wait time.Duration) BatchOption {
	return func(q *chainData) {
		q.wait = wait
	}
}

func OptNoWait() BatchOption {
	return func(q *chainData) {
		q.wait = -1
	}
}
