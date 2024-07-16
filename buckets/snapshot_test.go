package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"sync"
	"testing"
	"time"
)

func TestCleanupAfterSnapshot(t *testing.T) {
	wg := sync.WaitGroup{}
	cleanupInterval := 10
	b := NewBucketsRecorder(WithSnapshotOperation(CleanupAfterSnapshot(uint64(cleanupInterval))))

	for i := 0; i < cleanupInterval*2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.Record(duration.NewDuration(time.Duration(rand.Int64())))
			sp := recorder.RecordedSnapshot{}
			b.Snapshot(sp)
			assert.False(t, sp.IsEmpty())
		}()
	}
	wg.Wait()
	sp := recorder.RecordedSnapshot{}
	b.Snapshot(sp)
	assert.True(t, sp.IsEmpty())

}

func TestCleanupAndThenSwitch(t *testing.T) {
	wg := sync.WaitGroup{}
	cleanupInterval := 10
	b := NewBucketsRecorder(WithSnapshotOperation(CleanupAndThenSwitch(uint64(cleanupInterval), OnlySnapshot())))

	for i := 0; i < cleanupInterval*2; i++ {
		if i == cleanupInterval {
			wg.Wait()
			sp := recorder.RecordedSnapshot{}
			b.Snapshot(sp)
			assert.True(t, sp.IsEmpty())
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.Record(duration.NewDuration(time.Duration(rand.Int64())))
			sp := recorder.RecordedSnapshot{}
			b.Snapshot(sp)
			assert.False(t, sp.IsEmpty())
		}()
	}
	wg.Wait()
	sp := recorder.RecordedSnapshot{}
	b.Snapshot(sp)
	assert.False(t, sp.IsEmpty())
}
