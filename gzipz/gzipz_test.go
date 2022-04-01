package gzipz_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-bites/gzipz"
)

func TestCompress(t *testing.T) {
	buf := []byte(strings.Repeat("x", 1000))
	gBuf := gzipz.MustCompress(buf)
	require.Less(t, len(gBuf), len(buf))
	uBuf := gzipz.MustDecompress(gBuf)
	require.Equal(t, buf, uBuf)
}
