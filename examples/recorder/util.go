package utils

import (
	"fmt"
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"
	"time"
)

func InitData(r recorder.Recorder) {
	count := 200_000
	perInterval := (3 * time.Second).Nanoseconds() / int64(count)
	for i := 0; i < count; i++ {
		// 0.00s~3.00s
		r.Record(duration.NewDuration(time.Duration(int64(i) * perInterval)))
	}
}

func PrintReport(sp recorder.RecordedSnapshot) {
	mmax := sp.Max()
	mmin := sp.Min()
	mean := sp.Mean()
	p10 := sp.Percentile(10)
	p20 := sp.Percentile(20)
	p30 := sp.Percentile(30)
	p50 := sp.Percentile(50)
	p90 := sp.Percentile(90)
	p99 := sp.Percentile(99)
	p999 := sp.Percentile(99.9)
	str := fmt.Sprintf("max:%s min:%s mean:%s \np10:%s p20:%s p30:%s \np50:%s p90:%s p99:%s p999:%s",
		mmax.String(),
		mmin.String(),
		mean.String(),
		p10.String(),
		p20.String(),
		p30.String(),
		p50.String(),
		p90.String(),
		p99.String(),
		p999.String())
	fmt.Println(str)
}
