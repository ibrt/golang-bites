package enums_test

import (
	"testing"

	"github.com/ibrt/golang-bites/enumz/internal/testing/enums"

	"github.com/stretchr/testify/require"
)

func TestOutputMode(t *testing.T) {
	require.Equal(t, "batched", enums.OutputModeBatched.String())
	require.Equal(t, "streaming", enums.OutputModeStreaming.String())
	require.Equal(t, "batched_streaming", enums.OutputModeBatchedStreaming.String())

	require.True(t, enums.OutputModeBatched.Valid())
	require.True(t, enums.OutputModeStreaming.Valid())
	require.True(t, enums.OutputModeBatchedStreaming.Valid())
	require.False(t, enums.OutputMode("").Valid())
	require.False(t, enums.OutputMode("").Valid())
	require.False(t, enums.OutputMode("").Valid())
}
