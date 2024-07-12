# latency

[English](README.md)
## 简介
这是一个并发安全的用于统计延时的组件，采用分桶实现，可容纳足够多的时延数据。

## 特性
- [Duration][duration.Duration] 对官方Duration的拓展，这是一个带单位的时间，可以轻松的在各个时间单位下转换，以及转换成最佳的单位。
- [Bucket][buckets.bucketsRecorder] 分桶记录器，提供了多种预置的分桶策略，如自动分桶，固定单位分桶，自定义序列分桶。
- [Snapshot][recoder.RecordedSnapshot] 快照，支持任意时间获取当前统计的均值,最大最小值,百分位值。
- 高性能，针对GC和并发进行了优化，最大程度减少性能开销。
- 并发安全

## 快速开始
引入latency包
``` shell
go get github.com/jhue58/latency
```

使用默认提供的reporter
``` go
r := report.NewLateReporter(buckets.NewBucketsRecorder())
wg := sync.WaitGroup{}
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        start, end := r.Alloc()
        start()
        time.Sleep(time.Duration(rand.Intn(600)) * time.Millisecond) // 模拟耗时
        end()
    }
}
wg.Wait()
fmt.Println(r.Report())
```

直接使用recorder（推荐）
``` go
r := buckets.NewBucketsRecorder()
wg := sync.WaitGroup{}
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        r.Record(duration.NewDuartionWithUnit(float64(rand.Intn(600)),duration.Ms)) // 模拟耗时 0~600ms
    }
}
wg.Wait()
sp := recorder.RecordedSnapshot{}
r.Snapshot(sp)
mean:=sp.Mean()
fmt.Println("mean: ",mean.String())

```

## 进阶
更多进阶示例可查看[examples](examples)