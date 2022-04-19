package batchlog

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

const _UNCLEAN_KEY = "-"

func (a *alpchain) clean() {
	// clean current node
	if a.key == "-" {
		return
	}

	a.data = nil
	if len(a.nexts) > 0 {
		return
	}

	a.nexts = nil

	// clean previous node
	delete(a.pre.nexts, a.key)
	if a.pre.data == nil && len(a.pre.nexts) == 0 {
		a.pre.clean()
	}
}

type chainData struct {
	event *zerolog.Event

	count            int
	maxRelativeBatch int // option

	start   time.Time
	timeout time.Duration // option

	lastUpdated time.Time
	wait        time.Duration
}

func (c *chainData) needLogged() bool {
	now := time.Now()
	return c.count >= c.maxRelativeBatch || // reached max batch counts
		(c.timeout != -1 && c.start.Add(c.timeout).Before(now)) || // timeout
		(c.wait != -1 && c.lastUpdated.Add(c.wait).Before(now)) // extend
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
		gap: time.Second,
		root: alpchain{
			nexts: make(map[string]*alpchain, 8),
		},
	}
	for i := zerolog.DebugLevel; i <= zerolog.PanicLevel; i++ {
		batcher.root.nexts[i.String()] = &alpchain{
			pre: &batcher.root,
			key: _UNCLEAN_KEY,
		}
		batcher.root.key = _UNCLEAN_KEY
	}
	go batcher.logging()
}

func ChangeGapTime(dur time.Duration) {
	batcher.gap = dur
}

func (b *chainedBatcher) Batch(e *event, opts ...BatchOption) {
	if len(e.batchKeysA) == 0 {
		e.event.Send()
		return
	}

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
			node.nexts[v] = &alpchain{
				key: v,
				pre: node,
			}
		}

		if k == len(e.batchKeysA)-1 {
			value, _ := node.nexts[v]
			if value.data == nil {
				value.data = rawChainData(e.event)

				if len(e.logger.opts) > 0 {
					value.data.applyOpts(e.logger.opts...)
				}
				value.data.applyOpts(opts...)
				b.logged = append(b.logged, *value)
			}

			value.data.count++
			value.data.lastUpdated = time.Now()
			node.nexts[v] = value
		}

		node = node.nexts[v]
	}
}

func rawChainData(event *zerolog.Event) *chainData {
	return &chainData{
		event:            event,
		count:            0,
		start:            time.Now(),
		timeout:          10 * time.Second,
		maxRelativeBatch: 20,
		wait:             5 * time.Second,
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

			batchOut(b, chainData, i)
		}
		b.chainMtx.Unlock()
	}
}

// batchOut log the message out and clean the alpchain (nexts, data)
// and remove that element out of logger
func batchOut(b *chainedBatcher, c *chainData, i int) {
	c.event.Int("__repeat", c.count)
	c.event.Send()

	b.logged[i].clean()
	b.logged.RemoveUnor(i)
}
