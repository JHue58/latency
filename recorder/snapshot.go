package recorder

import (
	"github.com/jhue58/latency/duration"
)

type RecordedSnapshot map[duration.Duration]int64

// Mean returns the mean value of the current snapshot
func (s RecordedSnapshot) Mean() (res duration.Duration) {
	if s.IsEmpty() {
		return
	}
	var count int64
	var lateCount float64
	for late, c := range s {
		count += c
		late.ToNs()
		lateCount += late.Value() * float64(c)
	}
	mean := lateCount / float64(count)
	res = duration.NewDurationWithUnit(mean, duration.Ns)
	res.ToBestUnit()
	return res
}

// Max returns the max value of the current snapshot
func (s RecordedSnapshot) Max() (res duration.Duration) {
	if s.IsEmpty() {
		return
	}
	var latency duration.Durations
	for late, _ := range s {
		latency = append(latency, late)
	}
	latency.Sort()
	return latency[len(latency)-1]
}

// Min returns the min value of the current snapshot
func (s RecordedSnapshot) Min() (res duration.Duration) {
	if s.IsEmpty() {
		return
	}
	var latency duration.Durations
	for late, _ := range s {
		latency = append(latency, late)
	}
	latency.Sort()
	return latency[0]
}

// Percentile returns the percentile value of the current snapshot.
func (s RecordedSnapshot) Percentile(percent float64) (res duration.Duration) {
	if s.IsEmpty() {
		return
	}
	var latency duration.Durations
	var count int64
	for late, c := range s {
		count += c
		latency = append(latency, late)
	}
	latency.Sort()

	idx := int64(float64(count) * (percent / 100))
	var rge int64
	for _, late := range latency {
		rge += s[late]
		if idx <= rge {
			return late
		}
	}

	return
}

// Count returns the count of the current snapshot.
func (s RecordedSnapshot) Count() (count int64) {
	for _, c := range s {
		count += c
	}
	return
}

func (s RecordedSnapshot) IsEmpty() bool {
	return len(s) <= 0
}
