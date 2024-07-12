package main

import (
	"github.com/jhue58/latency/buckets"
	"github.com/jhue58/latency/duration"
	utils "github.com/jhue58/latency/examples/recorder"
	"github.com/jhue58/latency/recorder"
)

func main() {
	r := buckets.NewBucketsRecorder(buckets.WithBucketAllocator(buckets.UnitBucketAllocator(duration.Ms)))
	utils.InitData(r)
	sp := recorder.RecordedSnapshot{}
	r.Snapshot(sp)
	utils.PrintReport(sp)
}
