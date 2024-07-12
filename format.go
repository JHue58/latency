package latency

import (
	"fmt"
	"github.com/jhue58/latency/recorder"
)

// FormatFunc is a function that formats the snapshot
type FormatFunc func(sp recorder.RecordedSnapshot) string

// DefaultFormat contains max,min,mean,p50,p90,p99,p999
func DefaultFormat(sp recorder.RecordedSnapshot) string {
	mmax := sp.Max()
	mmin := sp.Min()
	mean := sp.Mean()
	p50 := sp.Percentile(50)
	p90 := sp.Percentile(90)
	p99 := sp.Percentile(99)
	p999 := sp.Percentile(99.9)
	str := fmt.Sprintf("max:%s min:%s mean:%s p50:%s p90:%s p99:%s p999:%s",
		mmax.String(),
		mmin.String(),
		mean.String(),
		p50.String(),
		p90.String(),
		p99.String(),
		p999.String(),
	)
	return str
}
