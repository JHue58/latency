package duration

import "time"

type TimeUnit int64

const (
	// None use to represent the zero value
	None = TimeUnit(0)
	// Ns Nanosecond(ns)
	Ns = TimeUnit(time.Nanosecond)
	// Us Microsecond(μs)
	Us = TimeUnit(time.Microsecond)
	// Ms Millisecond(ms)
	Ms = TimeUnit(time.Millisecond)
	// S Second(s)
	S = TimeUnit(time.Second)
	// M Minute(m)
	M = TimeUnit(time.Minute)
	// H Hour(h)
	H = TimeUnit(time.Hour)
)

var units = []TimeUnit{
	None, Ns, Us, Ms, S, M, H,
}

// valid check
func (u TimeUnit) valid() (ok bool) {
	for _, unit := range units {
		if unit == u {
			return true
		}
	}
	return
}
