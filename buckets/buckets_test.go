package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"runtime"
	"testing"
	"time"
)

func TestBucketsRecorder_Cleanup(t *testing.T) {
	b := NewBucketsRecorder()
	for i := 0; i < 100; i++ {
		b.Record(duration.NewDuration(time.Duration(rand.Int64())))
	}
	sp := recorder.RecordedSnapshot{}
	b.Snapshot(sp)
	assert.False(t, sp.IsEmpty())

	b.Cleanup()
	sp = recorder.RecordedSnapshot{}
	b.Snapshot(sp)
	assert.True(t, sp.IsEmpty())
}

func TestBucketsRecorder_Reset(t *testing.T) {
	b := NewBucketsRecorder()
	for i := 0; i < 100; i++ {
		b.Record(duration.NewDuration(time.Duration(rand.Int64())))
	}
	sp := recorder.RecordedSnapshot{}
	b.Snapshot(sp)
	assert.False(t, sp.IsEmpty())

	b.Reset()
	sp = recorder.RecordedSnapshot{}
	b.Snapshot(sp)
	assert.True(t, sp.IsEmpty())
}

func BenchmarkSnapshot(b *testing.B) {

	buckets := NewBucketsRecorder()
	for i := 0; i < 100_000_000; i++ {
		buckets.Record(duration.NewDuration(time.Duration(rand.Int64N(int64(time.Hour)))))
	}
	b.Run("Snapshot", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sp := recorder.RecordedSnapshot{}
			buckets.Snapshot(sp)
		}
	})
	b.Run("SnapshotParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			b.ResetTimer()
			for pb.Next() {
				sp := recorder.RecordedSnapshot{}
				buckets.Snapshot(sp)
			}

		})
	})

}

func BenchmarkRecord(b *testing.B) {
	buckets := NewBucketsRecorder()
	for i := 0; i < 100_000_000; i++ {
		buckets.Record(duration.NewDuration(time.Duration(rand.Int64N(int64(time.Hour)))))
	}
	b.Run("Record", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buckets.Record(duration.NewDuration(time.Duration(rand.Int64N(int64(time.Hour)))))
		}
	})
	b.Run("RecordParallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			b.ResetTimer()
			for pb.Next() {
				buckets.Record(duration.NewDuration(time.Duration(rand.Int64N(int64(time.Hour)))))
			}
		})
	})

}

func TestBucketsRecorder_MemUse(t *testing.T) {
	buckets := NewBucketsRecorder()
	runtime.GC()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	before := mem.HeapInuse
	t.Logf("Befor Mem Used:%d byte", before)
	for i := 0; i < 100_000_000; i++ {
		buckets.Record(duration.NewDuration(time.Duration(rand.Int64N(int64(time.Hour)))))
	}

	sp := recorder.RecordedSnapshot{}
	buckets.Snapshot(sp)
	t.Logf("Recorded count: %d", sp.Count())
	runtime.ReadMemStats(&mem)
	t.Logf("After Mem Used:%d byte", mem.HeapInuse)
	t.Logf("Mem Used:%d byte", mem.HeapInuse-before)
}
