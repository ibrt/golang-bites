package uint8z

import (
	"sort"
	"strconv"
)

const (
	// BitSize is the size in bits of this type.
	BitSize = 8
)

var (
	_ sort.Interface = Slice{}
)

// Ptr returns a pointer to the value.
func Ptr(v uint8) *uint8 {
	return &v
}

// PtrZeroToNil returns a pointer to the value, or nil if 0.
func PtrZeroToNil(v uint8) *uint8 {
	if v == 0 {
		return nil
	}
	return &v
}

// PtrDefToNil returns a pointer to the value, or nil if "def".
func PtrDefToNil(v uint8, def uint8) *uint8 {
	if v == def {
		return nil
	}
	return &v
}

// Val returns the pointer value, defaulting to zero if nil.
func Val(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

// ValDef returns the pointer value, defaulting to "def" if nil.
func ValDef(v *uint8, def uint8) uint8 {
	if v == nil {
		return def
	}
	return *v
}

// Parse parses a string as base 10 uint8.
func Parse(v string) (uint8, error) {
	p, err := strconv.ParseUint(v, 10, BitSize)
	if err != nil {
		return 0, err
	}
	return (uint8)(p), nil
}

// MustParse is like Parse but panics on error.
func MustParse(v string) uint8 {
	p, err := Parse(v)
	if err != nil {
		panic(err)
	}
	return p
}

// Slice is a slice of values.
type Slice []uint8

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
func SliceToMap(s []uint8) map[uint8]struct{} {
	m := make(map[uint8]struct{}, len(s))
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

// MapToSlice converts a map to slice.
func MapToSlice(m map[uint8]struct{}) []uint8 {
	s := make([]uint8, 0, len(m))
	for v := range m {
		s = append(s, v)
	}
	return s
}

// SafeIndex returns "s[i]" if possible, and 0 otherwise.
func SafeIndex(s []uint8, i int) uint8 {
	if s == nil || i < 0 || i >= len(s) {
		return 0
	}
	return s[i]
}

// SafeIndexDef returns "s[i]" if possible, and "def" otherwise.
func SafeIndexDef(s []uint8, i int, def uint8) uint8 {
	if s == nil || i < 0 || i >= len(s) {
		return def
	}
	return s[i]
}

// SafeIndexPtr returns "s[i]" if possible, and nil otherwise.
func SafeIndexPtr(s []uint8, i int) *uint8 {
	if s == nil || i < 0 || i >= len(s) {
		return nil
	}
	return Ptr(s[i])
}
