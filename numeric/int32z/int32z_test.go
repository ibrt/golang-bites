package int32z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/int32z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int32z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int32z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int32z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int32z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int32z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int32(0), int32z.Val(nil))
	require.Equal(t, int32(0), int32z.Val(int32z.Ptr(0)))
	require.Equal(t, int32(1), int32z.Val(int32z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int32(1), int32z.ValDef(nil, 1))
	require.Equal(t, int32(0), int32z.ValDef(int32z.Ptr(0), 1))
	require.Equal(t, int32(1), int32z.ValDef(int32z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int32z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int32(10), v)
	_, err = int32z.Parse("")
	require.Error(t, err)
	_, err = int32z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, int32(10), int32z.MustParse("10")) })
	require.Panics(t, func() { int32z.MustParse("") })
	require.Panics(t, func() { int32z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := int32z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, int32z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int32]struct{}{}, int32z.SliceToMap(nil))
	require.Equal(t, map[int32]struct{}{}, int32z.SliceToMap([]int32{}))
	require.Equal(t, map[int32]struct{}{1: {}}, int32z.SliceToMap([]int32{1}))
	require.Equal(t, map[int32]struct{}{1: {}, 2: {}}, int32z.SliceToMap([]int32{1, 2}))
	require.Equal(t, map[int32]struct{}{1: {}, 2: {}}, int32z.SliceToMap([]int32{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int32{}, int32z.MapToSlice(nil))
	require.Equal(t, []int32{}, int32z.MapToSlice(map[int32]struct{}{}))
	require.Equal(t, []int32{1}, int32z.MapToSlice(map[int32]struct{}{1: {}}))
	require.Equal(t, map[int32]struct{}{1: {}, 2: {}}, int32z.SliceToMap(int32z.MapToSlice(map[int32]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int32(0), int32z.SafeIndex(nil, 0))
	require.Equal(t, int32(0), int32z.SafeIndex(nil, 1))
	require.Equal(t, int32(0), int32z.SafeIndex(nil, -1))
	require.Equal(t, int32(0), int32z.SafeIndex([]int32{}, 0))
	require.Equal(t, int32(0), int32z.SafeIndex([]int32{}, 1))
	require.Equal(t, int32(0), int32z.SafeIndex([]int32{}, -1))
	require.Equal(t, int32(1), int32z.SafeIndex([]int32{1}, 0))
	require.Equal(t, int32(0), int32z.SafeIndex([]int32{1}, 1))
	require.Equal(t, int32(0), int32z.SafeIndex([]int32{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, int32(100), int32z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, int32(100), int32z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, int32(100), int32z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, int32(100), int32z.SafeIndexDef([]int32{}, 0, 100))
	require.Equal(t, int32(100), int32z.SafeIndexDef([]int32{}, 1, 100))
	require.Equal(t, int32(100), int32z.SafeIndexDef([]int32{}, -1, 100))
	require.Equal(t, int32(1), int32z.SafeIndexDef([]int32{1}, 0, 100))
	require.Equal(t, int32(100), int32z.SafeIndexDef([]int32{1}, 1, 100))
	require.Equal(t, int32(100), int32z.SafeIndexDef([]int32{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, int32z.SafeIndexPtr(nil, 0))
	require.Nil(t, int32z.SafeIndexPtr(nil, 1))
	require.Nil(t, int32z.SafeIndexPtr(nil, -1))
	require.Nil(t, int32z.SafeIndexPtr([]int32{}, 0))
	require.Nil(t, int32z.SafeIndexPtr([]int32{}, 1))
	require.Nil(t, int32z.SafeIndexPtr([]int32{}, -1))
	require.Equal(t, int32z.Ptr(1), int32z.SafeIndexPtr([]int32{1}, 0))
	require.Nil(t, int32z.SafeIndexPtr([]int32{1}, 1))
	require.Nil(t, int32z.SafeIndexPtr([]int32{1}, -1))
}
