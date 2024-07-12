package duration

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDuration_To(t *testing.T) {
	nanoV := float64(time.Hour)
	d := NewDurationWithUnit(nanoV, Ns)
	d.ToNs()
	assert.Equal(t, nanoV, d.v)
	d.ToUs()
	assert.Equal(t, nanoV/float64(time.Microsecond), d.v)
	d.ToMs()
	assert.Equal(t, nanoV/float64(time.Millisecond), d.v)
	d.ToS()
	assert.Equal(t, nanoV/float64(time.Second), d.v)
	d.ToM()
	assert.Equal(t, nanoV/float64(time.Minute), d.v)
	d.ToH()
	assert.Equal(t, nanoV/float64(time.Hour), d.v)
}

func TestDuration_ToBestUnit(t *testing.T) {
	d := NewDurationWithUnit(1000, Ns)
	d.ToBestUnit()
	assert.Equal(t, float64(1), d.v)
	assert.Equal(t, Us, d.TimeUnit)

	d = NewDurationWithUnit(1000, Us)
	d.ToBestUnit()
	assert.Equal(t, float64(1), d.v)
	assert.Equal(t, Ms, d.TimeUnit)

	d = NewDurationWithUnit(0.001, Us)
	d.ToBestUnit()
	assert.Equal(t, float64(1), d.v)
	assert.Equal(t, Ns, d.TimeUnit)

	d = NewDurationWithUnit(100, None)
	d.ToBestUnit()
	assert.Equal(t, float64(100), d.v)
	assert.Equal(t, None, d.TimeUnit)
}

func TestDuration_ValueWithUnit(t *testing.T) {
	v := float64(time.Hour)
	d := NewDurationWithUnit(v, Ns)
	assert.Equal(t, v/float64(time.Nanosecond), d.ValueWithUnit(Ns))
	assert.Equal(t, v/float64(time.Microsecond), d.ValueWithUnit(Us))
	assert.Equal(t, v/float64(time.Millisecond), d.ValueWithUnit(Ms))
	assert.Equal(t, v/float64(time.Second), d.ValueWithUnit(S))
	assert.Equal(t, v/float64(time.Minute), d.ValueWithUnit(M))
	assert.Equal(t, v/float64(time.Hour), d.ValueWithUnit(H))
}

func TestDurations_Sort(t *testing.T) {
	var sorted, unsorted Durations
	sorted = []Duration{
		NewDuration(0),
		NewDurationWithUnit(100, Ns),
		NewDurationWithUnit(200, Us),
		NewDurationWithUnit(3000, Ms),
		NewDurationWithUnit(400, S),
		NewDurationWithUnit(500, M),
	}
	unsorted = []Duration{
		NewDurationWithUnit(3000, Ms),
		NewDurationWithUnit(100, Ns),
		NewDurationWithUnit(500, M),
		NewDuration(0),
		NewDurationWithUnit(400, S),
		NewDurationWithUnit(200, Us),
	}
	unsorted.Sort()
	assert.Equal(t, sorted, unsorted)

}

func BenchmarkNewDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDuration(20 * time.Second)
	}
}

func BenchmarkNewDurationWithUnit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDurationWithUnit(200, Ms)
	}
}
