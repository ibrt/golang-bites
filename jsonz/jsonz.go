package jsonz

import (
	"bytes"
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

// MustIndent is like json.Indent but panics on error.
func MustIndent(src []byte, prefix, indent string) []byte {
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, src, prefix, indent); err != nil {
		panic(err)
	}
	return dst.Bytes()
}
