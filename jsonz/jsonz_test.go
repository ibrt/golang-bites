package jsonz_test

import (
	"testing"

	"github.com/ibrt/golang-bites/jsonz"

	"github.com/stretchr/testify/require"
)

func TestMustMarshal(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustMarshal(map[string]int{"a": 1})
		require.Equal(t, `{"a":1}`, string(buf))
	})

	require.PanicsWithError(t, `json: unsupported type: func()`, func() {
		jsonz.MustMarshal(func() {})
	})
}

func TestMustMarshalIndent(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustMarshalIndent(map[string]int{"a": 1}, "#", "  ")
		require.Equal(t, "{\n#  \"a\": 1\n#}", string(buf))
	})

	require.PanicsWithError(t, `json: unsupported type: func()`, func() {
		jsonz.MustMarshalIndent(func() {}, "#", "  ")
	})
}

func TestMustIndent(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustIndent([]byte(`{"a":1}`), "#", "  ")
		require.Equal(t, "{\n#  \"a\": 1\n#}", string(buf))
	})

	require.PanicsWithError(t, `invalid character 'b' looking for beginning of value`, func() {
		jsonz.MustIndent([]byte(`bad`), "#", "  ")
	})
}
