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

// MustMarshalString is like MustMarshal but returns a string.
func MustMarshalString(v interface{}) string {
	return string(MustMarshal(v))
}

// MustMarshalIndent is like json.MarshalIndent but panics on error.
func MustMarshalIndent(v interface{}, prefix, indent string) []byte {
	buf, err := json.MarshalIndent(v, prefix, indent)
	internal.MaybePanic(err)
	return buf
}

// MustMarshalIndentString is like MustMarshalIndent but returns a string.
func MustMarshalIndentString(v interface{}, prefix, indent string) string {
	return string(MustMarshalIndent(v, prefix, indent))
}

// MustMarshalIndentDefault is like MustMarshalIndent with prefix = "" and indent = "  ".
func MustMarshalIndentDefault(v interface{}) []byte {
	return MustMarshalIndent(v, "", "  ")
}

// MustMarshalIndentDefaultString is like MustMarshalIndentDefault with prefix = "" and indent = "  ".
func MustMarshalIndentDefaultString(v interface{}) string {
	return MustMarshalIndentString(v, "", "  ")
}

// MustIndent is like json.Indent but panics on error.
func MustIndent(src []byte, prefix, indent string) []byte {
	dst := &bytes.Buffer{}
	internal.MaybePanic(json.Indent(dst, src, prefix, indent))
	return dst.Bytes()
}

// MustIndentString is like MustIndent but returns a string.
func MustIndentString(src []byte, prefix, indent string) string {
	return string(MustIndent(src, prefix, indent))
}

// MustIndentDefault is like MustIndent with prefix = "" and indent = "  ".
func MustIndentDefault(src []byte) []byte {
	return MustIndent(src, "", "  ")
}

// MustIndentDefaultString is like MustIndentString with prefix = "" and indent = "  ".
func MustIndentDefaultString(src []byte) string {
	return MustIndentString(src, "", "  ")
}
