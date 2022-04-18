package logtest

import "time"

type BatchOption func(*chainData)

func (c *chainData) applyOpts(opts ...BatchOption) *chainData {
	for i := range opts {
		opts[i](c)
	}

	return c
}

func Timeout(timeout time.Duration) BatchOption {
	return func(q *chainData) {
		q.timeout = timeout
	}
}

func MaxBatch(maxBatch int) BatchOption {
	return func(q *chainData) {
		q.maxBatch = maxBatch
	}
}
