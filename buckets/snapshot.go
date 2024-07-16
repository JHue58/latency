package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sync"
	"sync/atomic"
)

// SnapshotOperation is a function that performs an operation get map snapshot.
type SnapshotOperation func(mp cmap.ConcurrentMap[duration.Duration, counter], sp recorder.RecordedSnapshot)

// Snapshot gets the current bucket statistics snapshot.
func (b *bucketsRecorder) Snapshot(sp recorder.RecordedSnapshot) {
	b.opt.snapshotFunc(b.mp, sp)
}

// CleanupAfterSnapshot returns a SnapshotOperation that performs cleanup
// after a specified number of operations (cleanupInterval).
func CleanupAfterSnapshot(cleanupInterval uint64) SnapshotOperation {
	if cleanupInterval <= 0 {
		cleanupInterval = 1
	}
	intervalI64 := int64(cleanupInterval)
	var count atomic.Int64
	return func(mp cmap.ConcurrentMap[duration.Duration, counter], sp recorder.RecordedSnapshot) {
	RETRY:
		added := count.Add(1)
		if added < intervalI64 {
			mp.IterCb(func(key duration.Duration, v counter) {
				sp[key] += v.v.Load()
			})
		} else if added == intervalI64 {
			var keys []duration.Duration
			count.Store(0)
			mp.IterCb(func(key duration.Duration, v counter) {
				keys = append(keys, key)
			})
			for _, key := range keys {
				v, ok := mp.Pop(key)
				if ok {
					sp[key] += v.v.Load()
					v.recycle()
				}
			}
		} else {
			goto RETRY
		}

	}
}

// CleanupAndThenSwitch returns a SnapshotOperation that performs cleanup
// after a specified number of operations (cleanupInterval), and then switches
// to the provided SnapshotOperation.
func CleanupAndThenSwitch(cleanupInterval uint64, to SnapshotOperation) SnapshotOperation {
	if cleanupInterval <= 0 {
		cleanupInterval = 1
	}
	intervalI64 := int64(cleanupInterval)
	var count atomic.Int64
	var cleaned atomic.Bool
	var cleanLock sync.RWMutex

	return func(mp cmap.ConcurrentMap[duration.Duration, counter], sp recorder.RecordedSnapshot) {
	RETRY:
		if cleaned.Load() {
			cleanLock.RLock()
			to(mp, sp)
			cleanLock.RUnlock()
			return
		}
		added := count.Add(1)
		if added < intervalI64 {
			mp.IterCb(func(key duration.Duration, v counter) {
				sp[key] += v.v.Load()
			})
		} else if added == intervalI64 {
			var keys []duration.Duration
			cleanLock.Lock()
			defer cleanLock.Unlock()
			cleaned.Store(true)
			mp.IterCb(func(key duration.Duration, v counter) {
				keys = append(keys, key)
			})
			for _, key := range keys {
				v, ok := mp.Pop(key)
				if ok {
					sp[key] += v.v.Load()
					v.recycle()
				}
			}
		} else {
			goto RETRY
		}

	}

}

// OnlySnapshot returns a SnapshotOperation.
func OnlySnapshot() SnapshotOperation {
	return func(mp cmap.ConcurrentMap[duration.Duration, counter], sp recorder.RecordedSnapshot) {
		mp.IterCb(func(key duration.Duration, v counter) {
			sp[key] += v.v.Load()
		})
	}
}