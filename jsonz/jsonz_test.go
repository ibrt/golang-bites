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

func TestMustMarshalString(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustMarshalString(map[string]int{"a": 1})
		require.Equal(t, `{"a":1}`, buf)
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

func TestMustMarshalIndentString(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustMarshalIndentString(map[string]int{"a": 1}, "#", "  ")
		require.Equal(t, "{\n#  \"a\": 1\n#}", buf)
	})

	require.PanicsWithError(t, `json: unsupported type: func()`, func() {
		jsonz.MustMarshalIndentString(func() {}, "#", "  ")
	})
}

func TestMustMarshalIndentDefault(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustMarshalIndentDefault(map[string]int{"a": 1})
		require.Equal(t, "{\n  \"a\": 1\n}", string(buf))
	})

	require.PanicsWithError(t, `json: unsupported type: func()`, func() {
		jsonz.MustMarshalIndentDefault(func() {})
	})
}

func TestMustMarshalIndentDefaultString(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustMarshalIndentDefaultString(map[string]int{"a": 1})
		require.Equal(t, "{\n  \"a\": 1\n}", buf)
	})

	require.PanicsWithError(t, `json: unsupported type: func()`, func() {
		jsonz.MustMarshalIndentDefaultString(func() {})
	})
}

func TestMustIndent(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustIndent([]byte(`{"a":1}`), "#", "  ")
		require.Equal(t, "{\n#  \"a\": 1\n#}", string(buf))
	})

	require.PanicsWithError(t, `invalid character 'b' looking for beginning of value`, func() {
		jsonz.MustIndentString([]byte(`bad`), "#", "  ")
	})
}

func TestMustIndentString(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustIndentString([]byte(`{"a":1}`), "#", "  ")
		require.Equal(t, "{\n#  \"a\": 1\n#}", buf)
	})

	require.PanicsWithError(t, `invalid character 'b' looking for beginning of value`, func() {
		jsonz.MustIndentString([]byte(`bad`), "#", "  ")
	})
}

func TestMustIndentDefault(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustIndentDefault([]byte(`{"a":1}`))
		require.Equal(t, "{\n  \"a\": 1\n}", string(buf))
	})

	require.PanicsWithError(t, `invalid character 'b' looking for beginning of value`, func() {
		jsonz.MustIndentDefault([]byte(`bad`))
	})
}

func TestMustIndentDefaultString(t *testing.T) {
	require.NotPanics(t, func() {
		buf := jsonz.MustIndentDefaultString([]byte(`{"a":1}`))
		require.Equal(t, "{\n  \"a\": 1\n}", buf)
	})

	require.PanicsWithError(t, `invalid character 'b' looking for beginning of value`, func() {
		jsonz.MustIndentDefaultString([]byte(`bad`))
	})
}
