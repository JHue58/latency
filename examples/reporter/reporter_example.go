package main

import (
	"fmt"
	"github.com/jhue58/latency"
	"github.com/jhue58/latency/buckets"
	"math/rand"
	"time"
)

func main() {
	r := latency.NewLateReporter(buckets.NewBucketsRecorder())
	timer := time.NewTimer(5 * time.Second)
	defer func() {
		timer.Stop()
		fmt.Println(r.Report())
	}()
	for {
		select {
		case <-timer.C:
			return
		default:
		}
		go func() {
			start, end := r.Alloc()
			start()
			time.Sleep(time.Duration(rand.Intn(600)) * time.Millisecond)
			end()
		}()
	}
}
