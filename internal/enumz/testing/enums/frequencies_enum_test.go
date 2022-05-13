package enums

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrequency(t *testing.T) {
	require.Equal(t, "single", FrequencySingle.String())
	require.Equal(t, "multi", FrequencyMulti.String())

	require.Equal(t, "Single", FrequencySingle.Label())
	require.Equal(t, "Multi", FrequencyMulti.Label())

	require.True(t, FrequencySingle.Valid())
	require.True(t, FrequencyMulti.Valid())
	require.False(t, Frequency("").Valid())
	require.False(t, Frequency("").Valid())
}
