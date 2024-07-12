package recorder

import "github.com/jhue58/latency/duration"

// Recorder to record the duration.
type Recorder interface {
	// Record the duration.
	Record(d duration.Duration)
	// Snapshot gets the current statistics snapshot.
	Snapshot(sp RecordedSnapshot)
}
