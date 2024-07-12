package main

import (
	"github.com/jhue58/latency/buckets"
	utils "github.com/jhue58/latency/examples/recorder"
	"github.com/jhue58/latency/recorder"
)

func main() {
	r := buckets.NewBucketsRecorder(buckets.WithBucketAllocator(buckets.BestBucketAllocator()))
	utils.InitData(r)
	sp := recorder.RecordedSnapshot{}
	r.Snapshot(sp)
	utils.PrintReport(sp)
}
