package urlz_test

import (
	"net/url"
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

func TestMustUpdate(t *testing.T) {
	require.NotPanics(t, func() {
		require.Equal(t,
			"https://test/test",
			urlz.MustUpdate("https://test", func(u *url.URL) {
				u.Path = "/test"
			}))
	})

	require.PanicsWithError(t, `parse "\b": net/url: invalid control character in URL`, func() {
		urlz.MustUpdate("\b", func(u *url.URL) {})
	})
}
