package jsonz

import (
	"bytes"
	"encoding/json"
	"github.com/ibrt/golang-bites/internal"
)

// MustMarshal is like json.Marshal but panics on error.
func MustMarshal(v interface{}) []byte {
	buf, err := json.Marshal(v)
	internal.MaybePanic(err)
	return buf
}

// MustMarshalIndent is like json.MarshalIndent but panics on error.
func MustMarshalIndent(v interface{}, prefix, indent string) []byte {
	buf, err := json.MarshalIndent(v, prefix, indent)
	internal.MaybePanic(err)
	return buf
}

// MustIndent is like json.Indent but panics on error.
func MustIndent(src []byte, prefix, indent string) []byte {
	dst := &bytes.Buffer{}
	internal.MaybePanic(json.Indent(dst, src, prefix, indent))
	return dst.Bytes()
}
