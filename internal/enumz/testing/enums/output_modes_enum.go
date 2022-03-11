package enums

// OutputMode describes the output modes enum.
type OutputMode string

// String implements the fmt.Stringer interface.
func (v OutputMode) String() string {
	return string(v)
}

// Valid validates the enum.
func (v OutputMode) Valid() bool {
	switch v {
	case
		OutputModeBatched,
		OutputModeStreaming,
		OutputModeBatchedStreaming:
		return true
	default:
		return false
	}
}

// Known enum values.
const (
	OutputModeBatched          OutputMode = "batched"
	OutputModeStreaming        OutputMode = "streaming"
	OutputModeBatchedStreaming OutputMode = "batched_streaming"
)

// OutputModes lists the known enum values.
var (
	OutputModes = []OutputMode{
		OutputModeBatched,
		OutputModeStreaming,
		OutputModeBatchedStreaming,
	}
)
