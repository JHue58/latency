package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sync"
)

// bucketsRecorder bucket statistics.
// All methods are goroutine-safe.
type bucketsRecorder struct {
	mp  cmap.ConcurrentMap[duration.Duration, counter]
	opt option

	rsLock sync.RWMutex
}

// NewBucketsRecorder creates a bucket statistics instance.
func NewBucketsRecorder(opts ...Option) recorder.Recorder {
	mp := cmap.NewWithCustomShardingFunction[duration.Duration, counter](func(key duration.Duration) uint32 {
		return uint32(key.Value())
	})
	b := bucketsRecorder{mp: mp, opt: defaultOption}
	for _, fn := range opts {
		fn(&b.opt)
	}
	return &b
}

// Record a duration, which will be allocated to the corresponding bucket based on BucketSelector.
func (b *bucketsRecorder) Record(d duration.Duration) {
	key := b.opt.allocator(d)
REDO:
	val, ok := b.mp.Get(key)
	if !ok {
		b.mp.SetIfAbsent(key, newCounter())
		goto REDO
	}
	val.v.Add(1)
}

func (b *bucketsRecorder) Reset() {
	b.rsLock.Lock()
	defer b.rsLock.Unlock()
	b.Cleanup()
	b.opt.snapshotOp.do = b.opt.snapshotOp.reset()
}

func (b *bucketsRecorder) Cleanup() {
	b.mp.Clear()
}
