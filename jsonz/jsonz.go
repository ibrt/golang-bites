package jsonz

import (
	"encoding/json"
)

// MustMarshal is like json.Marshal but panics on error.
func MustMarshal(v interface{}) []byte {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return buf
}

// MustMarshalIndent is like json.MarshalIndent but panics on error.
func MustMarshalIndent(v interface{}, prefix, indent string) []byte {
	buf, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return buf
}
