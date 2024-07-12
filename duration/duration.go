package duration

import (
	"fmt"
	"sort"
	"time"
)

// Durations is a slice of Duration.
type Durations []Duration

// Sort sorts the durations in ascending order.
func (d Durations) Sort() {

	sort.Slice(d, func(i, j int) bool {
		return d[i].ValueWithUnit(Ns) < d[j].ValueWithUnit(Ns)
	})

}

// Duration is a duration with a unit.
// Generated from time.Duration, it retains the original precision during unit conversions.
type Duration struct {
	TimeUnit
	v float64
}

func NewDuration(t time.Duration) Duration {
	return Duration{
		Ns,
		float64(t.Nanoseconds()),
	}
}

// NewDurationWithUnit creates a new Duration with the given value and unit.
//
//		Example:
//	 NewDurationWithUnit(6,duration.S) is 6.00s
//	 NewDurationWithUnit(6,duration.Ms) is 6.00ms
func NewDurationWithUnit(value float64, unit TimeUnit) Duration {
	if !unit.valid() || value < 0 {
		panic("invalid unit or the value less than zero")
	}
	return Duration{
		unit,
		value,
	}
}

// ToNs converts the duration to nanoseconds.
func (d *Duration) ToNs() {
	d.toUnit(Ns)
}

// ToUs converts the duration to microseconds.
func (d *Duration) ToUs() {
	d.toUnit(Us)
}

// ToMs converts the duration to milliseconds.
func (d *Duration) ToMs() {
	d.toUnit(Ms)
}

// ToS converts the duration to seconds.
func (d *Duration) ToS() {
	d.toUnit(S)
}

// ToM converts the duration to minutes.
func (d *Duration) ToM() {
	d.toUnit(M)
}

// ToH converts the duration to hours.
func (d *Duration) ToH() {
	d.toUnit(H)
}

// Value returns the value of the duration.
// Example:
//
//	4.96ms, the Value is 4.96
func (d *Duration) Value() float64 {
	return d.v
}

// ValueWithUnit returns the value of the duration with the given unit.
// Example:
//
//	4.96ms, the ValueWithUnit(Us) is 4960
func (d *Duration) ValueWithUnit(unit TimeUnit) float64 {
	if !unit.valid() {
		panic("invalid unit")
	}
	if d.TimeUnit == unit {
		return d.v
	}
	ratio := float64(d.TimeUnit) / float64(unit)
	return d.v * ratio

}

func (d *Duration) toUnit(unit TimeUnit) {
	if d.TimeUnit == unit {
		return
	}
	if unit == None {
		if d.v > 0 {
			unit = Ns
		} else {
			d.zero()
			return
		}
	}
	ratio := float64(d.TimeUnit) / float64(unit)
	d.v *= ratio
	d.TimeUnit = unit
}

func (d *Duration) zero() {
	d.TimeUnit = None
	d.v = 0
}

// String returns the format string of the duration like time.Duration.
func (d *Duration) String() (str string) {
	switch d.TimeUnit {
	case None:
		str = "None"
	case Ns:
		str = fmt.Sprintf("%.2fns", d.v)
	case Us:
		str = fmt.Sprintf("%.2fus", d.v)
	case Ms:
		str = fmt.Sprintf("%.2fms", d.v)
	case S:
		str = fmt.Sprintf("%.2fs", d.v)
	case M:
		str = fmt.Sprintf("%.2fm", d.v)
	case H:
		str = fmt.Sprintf("%.2fh", d.v)
	}
	return
}

// ToBestUnit converts the duration to the best unit.
// Example:
// 0.006ms -> 6.00us, 6000.00ms -> 6.00s
func (d *Duration) ToBestUnit() {
	length := len(units)
	nanoV := d.ValueWithUnit(Ns)
	for i := 0; i < length; i++ {
		if i < length-1 && nanoV < float64(units[i+1]) {
			d.toUnit(units[i])
			return
		}
	}
	d.toUnit(units[length-1])
}
