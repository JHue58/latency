package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"testing"
	"time"
)

func recordData(count int, r recorder.Recorder) {
	for i := 0; i < count; i++ {
		r.Record(duration.NewDuration(time.Duration(rand.Int64())))
	}

}

func isSnapshotEmpty(r recorder.Recorder) bool {
	sp := recorder.RecordedSnapshot{}
	r.Snapshot(sp)
	return sp.IsEmpty()
}

func TestCleanupAfterSnapshot(t *testing.T) {
	b := NewBucketsRecorder(WithSnapshotOperation(CleanupAfterSnapshot(2)))
	test := func() {
		recordData(10, b)
		assert.False(t, isSnapshotEmpty(b)) // count = 1
		assert.False(t, isSnapshotEmpty(b)) // count = 2 (cleanup)

		assert.True(t, isSnapshotEmpty(b)) // count = 1, if not Reset(),in next test() count will equal 2
	}
	test()
	b.Reset()
	test()
}

func TestCleanupAndThenSwitch(t *testing.T) {

	b := NewBucketsRecorder(WithSnapshotOperation(CleanupAndThenSwitch(1, CleanupAfterSnapshot(2))))
	test := func() {
		recordData(10, b)

		assert.False(t, isSnapshotEmpty(b))
		// switch

		assert.True(t, isSnapshotEmpty(b)) // count = 1

		recordData(10, b)
		assert.False(t, isSnapshotEmpty(b)) // count = 2 (cleanup)
		assert.True(t, isSnapshotEmpty(b))  // count = 1
		recordData(10, b)
		assert.False(t, isSnapshotEmpty(b)) // count = 2 (cleanup)
	}
	test()
	b.Reset()
	test()

}
