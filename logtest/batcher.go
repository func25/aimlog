package logtest

import (
	"sort"
	"time"
)

type alpchain struct {
	nexts map[string]alpchain
	data  *chainData
}

type chainData struct {
	count    int
	start    int64
	timeout  time.Duration
	maxBatch int
}

type batcher struct {
	root alpchain
}

func (b *batcher) Batch(e *event, opts ...BatchOption) {
	sort.Slice(e.batchKeysA, func(i, j int) bool {
		return e.batchKeysA[i] < e.batchKeysA[j]
	})

	if b.root.nexts == nil {
		b.root.nexts = make(map[string]alpchain)
	}
	node := b.root
	for k, v := range e.batchKeysA {
		if node.nexts == nil {
			node.nexts = make(map[string]alpchain, 1)
		}

		if _, ok := node.nexts[v]; !ok {
			node.nexts[v] = alpchain{}
		}

		if k == len(e.batchKeysA) {
			value, _ := node.nexts[v]
			if value.data == nil {
				value.data = &chainData{}
			}

			value.data.count++
			value.data.applyOpts(opts...)
			node.nexts[v] = value
		}
	}
}
