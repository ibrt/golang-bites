package stringz

import (
	"sort"
)

var (
	_ sort.Interface = Slice{}
)

// Ptr returns a pointer to the value.
func Ptr(v string) *string {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if "".
func PtrZeroToNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v string, def string) *string {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to "" if nil.
func Val(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *string, def string) string {
	if v == nil {
		return def
	}
	return *v
}

// Slice is a slice of values.
type Slice []string

// Len implements the sort.Interface interface.
func (s Slice) Len() int {
	return len(s)
}

// Less implements the sort.Interface interface.
func (s Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap implements the sort.Interface interface.
func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SliceToMap converts a slice to map.
func SliceToMap(s []string) map[string]struct{} {
	m := make(map[string]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[string]struct{}) []string {
	s := make([]string, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}

// SafeIndex returns "s[i]" if possible, and 0 otherwise.
func SafeIndex(s []string, i int) string {
	if s == nil || i < 0 || i >= len(s) {
		return ""
	}
	return s[i]
}

// SafeIndexDef returns "s[i]" if possible, and "def" otherwise.
func SafeIndexDef(s []string, i int, def string) string {
	if s == nil || i < 0 || i >= len(s) {
		return def
	}
	return s[i]
}

// SafeIndexPtr returns "s[i]" if possible, and nil otherwise.
func SafeIndexPtr(s []string, i int) *string {
	if s == nil || i < 0 || i >= len(s) {
		return nil
	}
	return Ptr(s[i])
}
