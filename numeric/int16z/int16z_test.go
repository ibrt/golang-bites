package int16z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/int16z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int16z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int16(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int16z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int16z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int16(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int16z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int16z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int16(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int16(0), int16z.Val(nil))
	require.Equal(t, int16(0), int16z.Val(int16z.Ptr(0)))
	require.Equal(t, int16(1), int16z.Val(int16z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int16(1), int16z.ValDef(nil, 1))
	require.Equal(t, int16(0), int16z.ValDef(int16z.Ptr(0), 1))
	require.Equal(t, int16(1), int16z.ValDef(int16z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int16z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int16(10), v)
	_, err = int16z.Parse("")
	require.Error(t, err)
	_, err = int16z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, int16(10), int16z.MustParse("10")) })
	require.Panics(t, func() { int16z.MustParse("") })
	require.Panics(t, func() { int16z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := int16z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, int16z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int16]struct{}{}, int16z.SliceToMap(nil))
	require.Equal(t, map[int16]struct{}{}, int16z.SliceToMap([]int16{}))
	require.Equal(t, map[int16]struct{}{1: {}}, int16z.SliceToMap([]int16{1}))
	require.Equal(t, map[int16]struct{}{1: {}, 2: {}}, int16z.SliceToMap([]int16{1, 2}))
	require.Equal(t, map[int16]struct{}{1: {}, 2: {}}, int16z.SliceToMap([]int16{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int16{}, int16z.MapToSlice(nil))
	require.Equal(t, []int16{}, int16z.MapToSlice(map[int16]struct{}{}))
	require.Equal(t, []int16{1}, int16z.MapToSlice(map[int16]struct{}{1: {}}))
	require.Equal(t, map[int16]struct{}{1: {}, 2: {}}, int16z.SliceToMap(int16z.MapToSlice(map[int16]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int16(0), int16z.SafeIndex(nil, 0))
	require.Equal(t, int16(0), int16z.SafeIndex(nil, 1))
	require.Equal(t, int16(0), int16z.SafeIndex(nil, -1))
	require.Equal(t, int16(0), int16z.SafeIndex([]int16{}, 0))
	require.Equal(t, int16(0), int16z.SafeIndex([]int16{}, 1))
	require.Equal(t, int16(0), int16z.SafeIndex([]int16{}, -1))
	require.Equal(t, int16(1), int16z.SafeIndex([]int16{1}, 0))
	require.Equal(t, int16(0), int16z.SafeIndex([]int16{1}, 1))
	require.Equal(t, int16(0), int16z.SafeIndex([]int16{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, int16(100), int16z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, int16(100), int16z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, int16(100), int16z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, int16(100), int16z.SafeIndexDef([]int16{}, 0, 100))
	require.Equal(t, int16(100), int16z.SafeIndexDef([]int16{}, 1, 100))
	require.Equal(t, int16(100), int16z.SafeIndexDef([]int16{}, -1, 100))
	require.Equal(t, int16(1), int16z.SafeIndexDef([]int16{1}, 0, 100))
	require.Equal(t, int16(100), int16z.SafeIndexDef([]int16{1}, 1, 100))
	require.Equal(t, int16(100), int16z.SafeIndexDef([]int16{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, int16z.SafeIndexPtr(nil, 0))
	require.Nil(t, int16z.SafeIndexPtr(nil, 1))
	require.Nil(t, int16z.SafeIndexPtr(nil, -1))
	require.Nil(t, int16z.SafeIndexPtr([]int16{}, 0))
	require.Nil(t, int16z.SafeIndexPtr([]int16{}, 1))
	require.Nil(t, int16z.SafeIndexPtr([]int16{}, -1))
	require.Equal(t, int16z.Ptr(1), int16z.SafeIndexPtr([]int16{1}, 0))
	require.Nil(t, int16z.SafeIndexPtr([]int16{1}, 1))
	require.Nil(t, int16z.SafeIndexPtr([]int16{1}, -1))
}
