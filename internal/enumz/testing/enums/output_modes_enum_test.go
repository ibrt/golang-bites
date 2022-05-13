package enums

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOutputMode(t *testing.T) {
	require.Equal(t, "batched", OutputModeBatched.String())
	require.Equal(t, "streaming", OutputModeStreaming.String())
	require.Equal(t, "batched_streaming", OutputModeBatchedStreaming.String())

	require.Equal(t, "Batched", OutputModeBatched.Label())
	require.Equal(t, "Streaming", OutputModeStreaming.Label())
	require.Equal(t, "BatchedStreaming", OutputModeBatchedStreaming.Label())

	require.True(t, OutputModeBatched.Valid())
	require.True(t, OutputModeStreaming.Valid())
	require.True(t, OutputModeBatchedStreaming.Valid())
	require.False(t, OutputMode("").Valid())
	require.False(t, OutputMode("").Valid())
	require.False(t, OutputMode("").Valid())
}
