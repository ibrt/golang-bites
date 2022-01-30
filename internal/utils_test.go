package internal_test

import (
	"fmt"
	"github.com/ibrt/golang-bites/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaybePanic(t *testing.T) {
	require.NotPanics(t, func() {
		internal.MaybePanic(nil)
	})

	require.PanicsWithError(t, "test", func() {
		internal.MaybePanic(fmt.Errorf("test"))
	})
}
