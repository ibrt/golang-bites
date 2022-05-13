package enums

// OutputMode describes the output modes enum.
type OutputMode string

// String implements the fmt.Stringer interface.
func (v OutputMode) String() string {
	return string(v)
}

// Label returns the enum value label.
func (v OutputMode) Label() string {
	switch v {
	case OutputModeBatched:
		return "Batched"
	case OutputModeStreaming:
		return "Streaming"
	case OutputModeBatchedStreaming:
		return "BatchedStreaming"
	default:
		return ""
	}
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
