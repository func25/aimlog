package logtest

import (
	"sort"
	"sync"
	"time"

	"github.com/func25/slicesol/slicesol"
	"github.com/rs/zerolog"
)

type alpchain struct {
	key   string
	nexts map[string]*alpchain
	pre   *alpchain
	data  *chainData
}

func (a *alpchain) clean() {
	a.data = nil
	if len(a.nexts) > 0 {
		return
	}

	a.nexts = nil
	delete(a.pre.nexts, a.key)
}

type chainData struct {
	event *zerolog.Event

	count    int
	start    time.Time
	timeout  time.Duration
	maxBatch int
}

func (c *chainData) needLogged() bool {
	// x := c.start.Add(c.timeout).Unix()
	// y := time.Now().Unix() - x
	// y = y + 1
	return c.count > c.maxBatch || c.start.Add(c.timeout).Before(time.Now())
}

type chainedBatcher struct {
	root     alpchain
	chainMtx sync.Mutex

	logged slicesol.Sliol[alpchain]
	stop   bool

	gap time.Duration
}

var batcher chainedBatcher

func init() {
	batcher = chainedBatcher{
		gap: 1 * time.Second,
		root: alpchain{
			nexts: make(map[string]*alpchain, 8),
		},
	}
	for i := zerolog.DebugLevel; i <= zerolog.PanicLevel; i++ {
		batcher.root.nexts[i.String()] = &alpchain{
			pre: &batcher.root,
		}
	}
	go batcher.logging()
}

func (b *chainedBatcher) Batch(e *event, opts ...BatchOption) {
	sort.Slice(e.batchKeysA, func(i, j int) bool {
		return e.batchKeysA[i] < e.batchKeysA[j]
	})

	if b.root.nexts == nil {
		b.root.nexts = make(map[string]*alpchain)
	}
	node := b.root.nexts[e.level.String()]

	b.chainMtx.Lock()
	defer b.chainMtx.Unlock()

	// then batch keys
	for k, v := range e.batchKeysA {
		if node.nexts == nil {
			node.nexts = make(map[string]*alpchain, 1)
		}

		if _, ok := node.nexts[v]; !ok {
			node.nexts[v] = &alpchain{}
		}

		if k == len(e.batchKeysA)-1 {
			value, _ := node.nexts[v]
			if value.data == nil {
				value.data = rawChainData(e.event)
				value.pre = node
				value.key = v

				value.data.applyOpts(opts...)
				b.logged = append(b.logged, *value)
			}

			value.data.count++
			node.nexts[v] = value
		}

		node = node.nexts[v]
	}
}

func rawChainData(event *zerolog.Event) *chainData {
	return &chainData{
		event:    event,
		count:    0,
		start:    time.Now(),
		timeout:  2 * time.Second,
		maxBatch: 20,
	}
}

//functask: goroutines slice b.logged into several pieces
func (b *chainedBatcher) logging() {
	for ; !b.stop; time.Sleep(b.gap) {
		chainData := &chainData{}
		loggedLen := len(b.logged)

		b.chainMtx.Lock()
		for i := loggedLen - 1; i >= 0; i-- {
			chainData = b.logged[i].data
			// if b.logged[i].data == nil {
			// }

			if !chainData.needLogged() {
				continue
			}

			chainData.event.Int("__repeat", b.logged[i].data.count)
			chainData.event.Send()
			b.logged[i].clean()
			b.logged.RemoveUnor(i)
		}
		b.chainMtx.Unlock()
	}
}
