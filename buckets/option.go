package buckets

type option struct {
	allocator          BucketAllocator
	clearAfterSnapshot bool
}

// Option bucket configuration.
type Option func(o *option)

// WithBucketAllocator sets the BucketAllocator.
// the Record method will allocate buckets based on the result of BucketAllocator (default is BestBucketAllocator).
func WithBucketAllocator(sf BucketAllocator) Option {
	return func(o *option) {
		o.allocator = sf
	}
}

// WithClearAfterSnapshot set to clear the buckets after Snapshot.
func WithClearAfterSnapshot() Option {
	return func(o *option) {
		o.clearAfterSnapshot = true
	}
}

var defaultOption = option{
	allocator:          BestBucketAllocator(),
	clearAfterSnapshot: false,
}
