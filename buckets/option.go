package buckets

type option struct {
	allocator  BucketAllocator
	snapshotOp SnapshotOperation
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

// WithSnapshotOperation sets the SnapshotOperation (default is OnlySnapshot).
// SnapshotOperation is provided by the buckets package.
//
//	Example:
//
// WithSnapshotOperation(CleanupAfterSnapshot(6))
func WithSnapshotOperation(op SnapshotOperation) Option {
	return func(o *option) {
		o.snapshotOp = op
	}
}

var defaultOption = option{
	allocator:  BestBucketAllocator(),
	snapshotOp: OnlySnapshot(),
}
