package urlz

import (
	"net/url"
)

// MustParse is like url.Parse but panics on error.
func MustParse(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}
