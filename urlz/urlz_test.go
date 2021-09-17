package urlz_test

import (
	"testing"

	"github.com/ibrt/golang-bites/urlz"

	"github.com/stretchr/testify/require"
)

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() {
		u := urlz.MustParse("https://test")
		require.Equal(t, "https://test", u.String())
	})

	require.PanicsWithError(t, `parse "\b": net/url: invalid control character in URL`, func() {
		urlz.MustParse("\b")
	})
}
