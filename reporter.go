package latency

import (
	"github.com/jhue58/latency/duration"
	"github.com/jhue58/latency/recorder"

	"time"
)

// LateReporter is a reporter that records the latency.
type LateReporter struct {
	r   recorder.Recorder
	opt option
}

func NewLateReporter(r recorder.Recorder, opts ...Option) *LateReporter {
	res := &LateReporter{
		r:   r,
		opt: defaultOption,
	}
	for _, opt := range opts {
		opt(&res.opt)
	}
	return res
}

// Alloc allocates a start and end function, it's disposable.
//
//	Example:
//	start,end:=r.Alloc()
//	start()
//	// do something
//	end()
func (r *LateReporter) Alloc() (start func(), end func() (late duration.Duration)) {
	var s time.Time
	start = func() {
		s = time.Now()
	}
	end = func() duration.Duration {
		late := duration.NewDuration(time.Since(s))
		r.r.Record(late)
		return late
	}
	return
}

// Report reports the latency.
func (r *LateReporter) Report() string {
	sp := recorder.RecordedSnapshot{}
	r.r.Snapshot(sp)
	return r.opt.format(sp)
}
