package intz_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/intz"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := intz.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := intz.PtrZeroToNil(0)
	require.Nil(t, p)
	p = intz.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := intz.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = intz.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int(0), intz.Val(nil))
	require.Equal(t, int(0), intz.Val(intz.Ptr(0)))
	require.Equal(t, int(1), intz.Val(intz.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int(1), intz.ValDef(nil, 1))
	require.Equal(t, int(0), intz.ValDef(intz.Ptr(0), 1))
	require.Equal(t, int(1), intz.ValDef(intz.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := intz.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int(10), v)
	_, err = intz.Parse("")
	require.Error(t, err)
	_, err = intz.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, int(10), intz.MustParse("10")) })
	require.Panics(t, func() { intz.MustParse("") })
	require.Panics(t, func() { intz.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := intz.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, intz.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int]struct{}{}, intz.SliceToMap(nil))
	require.Equal(t, map[int]struct{}{}, intz.SliceToMap([]int{}))
	require.Equal(t, map[int]struct{}{1: {}}, intz.SliceToMap([]int{1}))
	require.Equal(t, map[int]struct{}{1: {}, 2: {}}, intz.SliceToMap([]int{1, 2}))
	require.Equal(t, map[int]struct{}{1: {}, 2: {}}, intz.SliceToMap([]int{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int{}, intz.MapToSlice(nil))
	require.Equal(t, []int{}, intz.MapToSlice(map[int]struct{}{}))
	require.Equal(t, []int{1}, intz.MapToSlice(map[int]struct{}{1: {}}))
	require.Equal(t, map[int]struct{}{1: {}, 2: {}}, intz.SliceToMap(intz.MapToSlice(map[int]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int(0), intz.SafeIndex(nil, 0))
	require.Equal(t, int(0), intz.SafeIndex(nil, 1))
	require.Equal(t, int(0), intz.SafeIndex(nil, -1))
	require.Equal(t, int(0), intz.SafeIndex([]int{}, 0))
	require.Equal(t, int(0), intz.SafeIndex([]int{}, 1))
	require.Equal(t, int(0), intz.SafeIndex([]int{}, -1))
	require.Equal(t, int(1), intz.SafeIndex([]int{1}, 0))
	require.Equal(t, int(0), intz.SafeIndex([]int{1}, 1))
	require.Equal(t, int(0), intz.SafeIndex([]int{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, int(100), intz.SafeIndexDef(nil, 0, 100))
	require.Equal(t, int(100), intz.SafeIndexDef(nil, 1, 100))
	require.Equal(t, int(100), intz.SafeIndexDef(nil, -1, 100))
	require.Equal(t, int(100), intz.SafeIndexDef([]int{}, 0, 100))
	require.Equal(t, int(100), intz.SafeIndexDef([]int{}, 1, 100))
	require.Equal(t, int(100), intz.SafeIndexDef([]int{}, -1, 100))
	require.Equal(t, int(1), intz.SafeIndexDef([]int{1}, 0, 100))
	require.Equal(t, int(100), intz.SafeIndexDef([]int{1}, 1, 100))
	require.Equal(t, int(100), intz.SafeIndexDef([]int{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, intz.SafeIndexPtr(nil, 0))
	require.Nil(t, intz.SafeIndexPtr(nil, 1))
	require.Nil(t, intz.SafeIndexPtr(nil, -1))
	require.Nil(t, intz.SafeIndexPtr([]int{}, 0))
	require.Nil(t, intz.SafeIndexPtr([]int{}, 1))
	require.Nil(t, intz.SafeIndexPtr([]int{}, -1))
	require.Equal(t, intz.Ptr(1), intz.SafeIndexPtr([]int{1}, 0))
	require.Nil(t, intz.SafeIndexPtr([]int{1}, 1))
	require.Nil(t, intz.SafeIndexPtr([]int{1}, -1))
}
