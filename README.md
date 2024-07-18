
# latency

[中文](README_CN.md)

## Introduction

This is a Goroutine-safe component for latency statistics, implemented using buckets to accommodate a large amount of latency data.

## Features
- [Duration][duration.Duration]: Extension of the `time.Duration`, representing time with units, allowing easy conversion between different time units and optimizing for the best unit.
- [Bucket][buckets.bucketsRecorder]: Bucket recorder offering various bucket allocation methods, such as automatic allocator, fixed unit allocator, and custom sequence allocator.
- [Snapshot][recorder.RecordedSnapshot]: Snapshot feature supporting retrieval of current statistical averages, maximums, minimums, and percentile values at any time.
- High performance: with optimizations for GC and concurrency, minimizing performance overhead to the maximum extent. [Benchmark][Benchmark]
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

## Benchmark

The bucket allocation method is the automatic allocator. [BenchmarkCode][BenchmarkCode]

### Random Record, Preloading 100 Million Data in the Bucket
```
goos: linux
goarch: amd64
pkg: github.com/jhue58/latency/buckets
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkRecord
BenchmarkRecord/Record
BenchmarkRecord/Record-16               27292686                46.07 ns/op
BenchmarkRecord/RecordParallel
BenchmarkRecord/RecordParallel-16       100000000               10.56 ns/op
```

### Snapshot of 100 Million Data
```
goos: linux
goarch: amd64
pkg: github.com/jhue58/latency/buckets
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkSnapshot
BenchmarkSnapshot/Snapshot
BenchmarkSnapshot/Snapshot-16               8595            122031 ns/op
BenchmarkSnapshot/SnapshotParallel
BenchmarkSnapshot/SnapshotParallel-16      21300             53997 ns/op
```

### Memory Usage for Random Record of 100 Million Data
```
=== RUN   TestBucketsRecorder_MemUse
    buckets_test.go:97: Befor Mem Used:811008 byte
    buckets_test.go:104: Recorded count: 100000000
    buckets_test.go:106: After Mem Used:1204224 byte
    buckets_test.go:107: Mem Used:393216 byte
```


## More
For more advanced examples, refer to [examples][examples].


[duration.Duration]: duration/duration.go
[buckets.bucketsRecorder]: buckets/buckets.go
[recorder.RecordedSnapshot]: recorder/snapshot.go
[Benchmark]: #Benchmark
[BenchmarkCode]: buckets/buckets_test.go
[examples]: examples