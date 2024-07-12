package buckets

import (
	"github.com/jhue58/latency/duration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeriesAllocator(t *testing.T) {
	allocator := SeriesAllocator(
		duration.NewDurationWithUnit(50, duration.Ms),
		duration.NewDurationWithUnit(100, duration.Ms),
		duration.NewDurationWithUnit(150, duration.Ms),
		duration.NewDurationWithUnit(200, duration.Ms),
		duration.NewDurationWithUnit(500, duration.Ms),
		duration.NewDurationWithUnit(1, duration.S),
	)
	target := allocator(duration.NewDurationWithUnit(749, duration.Ms))
	assert.Equal(t, duration.NewDurationWithUnit(500, duration.Ms), target)
	target = allocator(duration.NewDurationWithUnit(800, duration.Ms))
	assert.Equal(t, duration.NewDurationWithUnit(1, duration.S), target)
}
