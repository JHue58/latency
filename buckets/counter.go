package buckets

import (
	"sync"
	"sync/atomic"
)

var int64Pool = sync.Pool{
	New: func() any {
		return new(atomic.Int64)
	},
}

// counter used for bucket counting.
type counter struct {
	v *atomic.Int64
}

func newCounter() counter {
	return counter{
		v: int64Pool.Get().(*atomic.Int64),
	}
}

func (c *counter) recycle() {
	c.v.Store(0)
	int64Pool.Put(c.v)
}
