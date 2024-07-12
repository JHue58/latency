package main

import (
	"github.com/jhue58/latency/buckets"
	"github.com/jhue58/latency/duration"
	utils "github.com/jhue58/latency/examples/recorder"
	"github.com/jhue58/latency/recorder"
)

func main() {
	series := []duration.Duration{
		duration.NewDurationWithUnit(10, duration.Ms),
		duration.NewDurationWithUnit(50, duration.Ms),
		duration.NewDurationWithUnit(100, duration.Ms),
		duration.NewDurationWithUnit(150, duration.Ms),
		duration.NewDurationWithUnit(200, duration.Ms),
		duration.NewDurationWithUnit(300, duration.Ms),
		duration.NewDurationWithUnit(400, duration.Ms),
		duration.NewDurationWithUnit(500, duration.Ms),
		duration.NewDurationWithUnit(1, duration.S),
		duration.NewDurationWithUnit(1.5, duration.S),
		duration.NewDurationWithUnit(2, duration.S),
	}

	r := buckets.NewBucketsRecorder(buckets.WithBucketAllocator(buckets.SeriesAllocator(series...)))
	utils.InitData(r)
	sp := recorder.RecordedSnapshot{}
	r.Snapshot(sp)
	utils.PrintReport(sp)
}
