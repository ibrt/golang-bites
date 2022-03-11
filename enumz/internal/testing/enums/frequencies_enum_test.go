package enums_test

import (
	"testing"

	"github.com/ibrt/golang-bites/enumz/internal/testing/enums"

	"github.com/stretchr/testify/require"
)

func TestFrequency(t *testing.T) {
	require.Equal(t, "single", enums.FrequencySingle.String())
	require.Equal(t, "multi", enums.FrequencyMulti.String())

	require.True(t, enums.FrequencySingle.Valid())
	require.True(t, enums.FrequencyMulti.Valid())
	require.False(t, enums.Frequency("").Valid())
	require.False(t, enums.Frequency("").Valid())
}
