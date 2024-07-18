package recorder

import "github.com/jhue58/latency/duration"

// Recorder to record the duration.
type Recorder interface {
	// Record the duration.
	Record(d duration.Duration)
	// Snapshot gets the current statistics snapshot.
	Snapshot(sp RecordedSnapshot)
	// Cleanup the recorder, only clean up all data in the recorder.
	Cleanup()
	// Reset the recorder, clean up all data and reset the state of the recorder,
	// making it as if it was just created.
	Reset()
}
