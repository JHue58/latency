
# latency

[中文](README_CN.md)

## Introduction

This is a Goroutine-safe component for latency statistics, implemented using buckets to accommodate a large amount of latency data.

## Features
- [Duration][duration.Duration]: Extension of the official `time.Duration`, representing time with units, allowing easy conversion between different time units and optimizing for the best unit.
- [Bucket][buckets.bucketsRecorder]: Bucket recorder offering various preset bucketing strategies such as automatic bucketing, fixed unit bucketing, and custom sequence bucketing.
- [Snapshot][recoder.RecordedSnapshot]: Snapshot feature supporting retrieval of current statistical averages, maximums, minimums, and percentile values at any time.
- High performance with optimizations for garbage collection (GC) and concurrency, minimizing performance overhead to the maximum extent.
- Goroutine-safe.

## Quick Start
Import the latency package:
``` shell
go get github.com/jhue58/latency
```

Use the provided default reporter:
``` go
r := report.NewLateReporter(buckets.NewBucketsRecorder())
wg := sync.WaitGroup{}
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        start, end := r.Alloc()
        start()
        time.Sleep(time.Duration(rand.Intn(600)) * time.Millisecond) // Simulate time-consuming operation
        end()
    }()
}
wg.Wait()
fmt.Println(r.Report())
```

Directly use the recorder (recommended):
``` go
r := buckets.NewBucketsRecorder()
wg := sync.WaitGroup{}
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        r.Record(duration.NewDuartionWithUnit(float64(rand.Intn(600)),duration.Ms)) // Simulate latency 0~600ms
    }()
}
wg.Wait()
sp := recorder.RecordedSnapshot{}
r.Snapshot(sp)
mean := sp.Mean()
fmt.Println("mean: ", mean.String())
```

## More
For more advanced examples, refer to [examples](examples).
