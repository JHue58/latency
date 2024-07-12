package buckets

import (
	"github.com/jhue58/latency/duration"
	"math"
	"sort"
)

// BucketAllocator bucket allocator.
type BucketAllocator func(d duration.Duration) (target duration.Duration)

// UnitBucketAllocator allocates buckets with a specific duration.TimeUnit granularity using math.Round.
//
//	Example:
//	UnitBucketAllocator(duration.Ms)
//	3.96ms->4.00ms, 4.01ms->4.00ms, 100.96μs->0.00ms
func UnitBucketAllocator(unit duration.TimeUnit) BucketAllocator {
	return func(d duration.Duration) (target duration.Duration) {
		v := math.Round(d.ValueWithUnit(unit))
		return duration.NewDurationWithUnit(v, unit)
	}
}

// BestBucketAllocator allocates buckets with the best duration.TimeUnit granularity using math.Round.
//
//	Example:
//	3.96ms->3.00ms, 4.01ms->4.00ms, 100.34μs->100μs, 1.24s->1.00s, 4000.00ms -> 4.00s
func BestBucketAllocator() BucketAllocator {
	return func(d duration.Duration) (target duration.Duration) {
		d.ToBestUnit()
		v := math.Round(d.Value())
		return duration.NewDurationWithUnit(v, d.TimeUnit)
	}
}

// SeriesAllocator allocates buckets to the closest in the given series.
//
//	Example:
//	series is [50ms, 100ms, 150ms, 200ms, 500ms, 1s]
//	60.96ms -> 50ms, 80.96ms -> 100ms, 160.96ms -> 150ms
func SeriesAllocator(series ...duration.Duration) BucketAllocator {
	ds := duration.Durations(series)
	ds.Sort()

	var nanoValues []float64
	for i := 0; i < len(ds); i++ {
		nanoValues = append(nanoValues, ds[i].ValueWithUnit(duration.Ns))
	}
	return func(d duration.Duration) (target duration.Duration) {
		nanoValueDst := d.ValueWithUnit(duration.Ns)
		idx := sort.Search(len(nanoValues), func(i int) bool {
			return nanoValues[i] >= nanoValueDst
		})
		if idx == 0 {
			return ds[0]
		}
		if idx == len(ds) {
			return ds[len(ds)-1]
		}
		before := nanoValues[idx-1]
		after := nanoValues[idx]
		if math.Abs(nanoValueDst-before) < math.Abs(nanoValueDst-after) {
			return ds[idx-1]
		}
		return ds[idx]
	}
}
