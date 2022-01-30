package urlz

import (
	"github.com/ibrt/golang-bites/internal"
	"net/url"
)

// MustParse is like url.Parse but panics on error.
func MustParse(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	internal.MaybePanic(err)
	return u
}

// MustUpdate parses the given URL, calls f to allow mutations, and then converts the URL back to string.
func MustUpdate(rawURL string, f func(*url.URL)) string {
	u := MustParse(rawURL)
	f(u)
	return u.String()
}
