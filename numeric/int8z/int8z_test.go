package int8z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/int8z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int8z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int8(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int8z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int8z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int8(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int8z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int8z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int8(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int8(0), int8z.Val(nil))
	require.Equal(t, int8(0), int8z.Val(int8z.Ptr(0)))
	require.Equal(t, int8(1), int8z.Val(int8z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int8(1), int8z.ValDef(nil, 1))
	require.Equal(t, int8(0), int8z.ValDef(int8z.Ptr(0), 1))
	require.Equal(t, int8(1), int8z.ValDef(int8z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int8z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int8(10), v)
	_, err = int8z.Parse("")
	require.Error(t, err)
	_, err = int8z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, int8(10), int8z.MustParse("10")) })
	require.Panics(t, func() { int8z.MustParse("") })
	require.Panics(t, func() { int8z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := int8z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, int8z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int8]struct{}{}, int8z.SliceToMap(nil))
	require.Equal(t, map[int8]struct{}{}, int8z.SliceToMap([]int8{}))
	require.Equal(t, map[int8]struct{}{1: {}}, int8z.SliceToMap([]int8{1}))
	require.Equal(t, map[int8]struct{}{1: {}, 2: {}}, int8z.SliceToMap([]int8{1, 2}))
	require.Equal(t, map[int8]struct{}{1: {}, 2: {}}, int8z.SliceToMap([]int8{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int8{}, int8z.MapToSlice(nil))
	require.Equal(t, []int8{}, int8z.MapToSlice(map[int8]struct{}{}))
	require.Equal(t, []int8{1}, int8z.MapToSlice(map[int8]struct{}{1: {}}))
	require.Equal(t, map[int8]struct{}{1: {}, 2: {}}, int8z.SliceToMap(int8z.MapToSlice(map[int8]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int8(0), int8z.SafeIndex(nil, 0))
	require.Equal(t, int8(0), int8z.SafeIndex(nil, 1))
	require.Equal(t, int8(0), int8z.SafeIndex(nil, -1))
	require.Equal(t, int8(0), int8z.SafeIndex([]int8{}, 0))
	require.Equal(t, int8(0), int8z.SafeIndex([]int8{}, 1))
	require.Equal(t, int8(0), int8z.SafeIndex([]int8{}, -1))
	require.Equal(t, int8(1), int8z.SafeIndex([]int8{1}, 0))
	require.Equal(t, int8(0), int8z.SafeIndex([]int8{1}, 1))
	require.Equal(t, int8(0), int8z.SafeIndex([]int8{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, int8(100), int8z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, int8(100), int8z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, int8(100), int8z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, int8(100), int8z.SafeIndexDef([]int8{}, 0, 100))
	require.Equal(t, int8(100), int8z.SafeIndexDef([]int8{}, 1, 100))
	require.Equal(t, int8(100), int8z.SafeIndexDef([]int8{}, -1, 100))
	require.Equal(t, int8(1), int8z.SafeIndexDef([]int8{1}, 0, 100))
	require.Equal(t, int8(100), int8z.SafeIndexDef([]int8{1}, 1, 100))
	require.Equal(t, int8(100), int8z.SafeIndexDef([]int8{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, int8z.SafeIndexPtr(nil, 0))
	require.Nil(t, int8z.SafeIndexPtr(nil, 1))
	require.Nil(t, int8z.SafeIndexPtr(nil, -1))
	require.Nil(t, int8z.SafeIndexPtr([]int8{}, 0))
	require.Nil(t, int8z.SafeIndexPtr([]int8{}, 1))
	require.Nil(t, int8z.SafeIndexPtr([]int8{}, -1))
	require.Equal(t, int8z.Ptr(1), int8z.SafeIndexPtr([]int8{1}, 0))
	require.Nil(t, int8z.SafeIndexPtr([]int8{1}, 1))
	require.Nil(t, int8z.SafeIndexPtr([]int8{1}, -1))
}
