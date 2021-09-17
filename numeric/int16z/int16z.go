package int16z

import (
	"sort"
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 16
)

var (
	_ sort.Interface = Slice{}
)

// Ptr returns a pointer to the value.
func Ptr(v int16) *int16 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v int16) *int16 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v int16, def int16) *int16 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *int16, def int16) int16 {
	if v == nil {
		return def
	}
	return *v
}

// Parse parses a string as base 10 int16.
func Parse(v string) (int16, error) {
	p, err := strconv.ParseInt(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (int16)(p), nil
}

// MustParse is like Parse but panics on error.
func MustParse(v string) int16 {
	p, err := Parse(v)
	if err != nil {
		panic(err)
	}
	return p
}

// Slice is a slice of values.
type Slice []int16

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
func SliceToMap(s []int16) map[int16]struct{} {
	m := make(map[int16]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[int16]struct{}) []int16 {
	s := make([]int16, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}

// SafeIndex returns "s[i]" if possible, and 0 otherwise.
func SafeIndex(s []int16, i int) int16 {
	if s == nil || i < 0 || i >= len(s) {
		return 0
	}
	return s[i]
}

// SafeIndexDef returns "s[i]" if possible, and "def" otherwise.
func SafeIndexDef(s []int16, i int, def int16) int16 {
	if s == nil || i < 0 || i >= len(s) {
		return def
	}
	return s[i]
}

// SafeIndexPtr returns "s[i]" if possible, and nil otherwise.
func SafeIndexPtr(s []int16, i int) *int16 {
	if s == nil || i < 0 || i >= len(s) {
		return nil
	}
	return Ptr(s[i])
}
