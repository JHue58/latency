package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	cmap "github.com/orcaman/concurrent-map/v2"
)

// bucketsRecorder bucket statistics.
// All methods are goroutine-safe.
type bucketsRecorder struct {
	mp  cmap.ConcurrentMap[duration.Duration, counter]
	opt option
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

// Snapshot gets the current bucket statistics snapshot.
func (b *bucketsRecorder) Snapshot(sp recorder.RecordedSnapshot) {

	if b.opt.clearAfterSnapshot {
		var keys []duration.Duration
		b.mp.IterCb(func(key duration.Duration, v counter) {
			keys = append(keys, key)
		})
		for _, key := range keys {
			v, ok := b.mp.Pop(key)
			if ok {
				sp[key] += v.v.Load()
				v.recycle()
			}
		}
	} else {
		b.mp.IterCb(func(key duration.Duration, v counter) {
			sp[key] += v.v.Load()
		})
	}

}
