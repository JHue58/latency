package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	"testing"
	"time"
)

func BenchmarkSnapshot(b *testing.B) {

	buckets := NewBucketsRecorder()
	for i := 0; i < 100_000; i++ {
		buckets.Record(duration.NewDuration(time.Duration(i + 1)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sp := recorder.RecordedSnapshot{}
		buckets.Snapshot(sp)
	}
}

func BenchmarkRecord(b *testing.B) {
	buckets := NewBucketsRecorder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buckets.Record(duration.NewDuration(time.Duration(i + 1)))
	}

}
