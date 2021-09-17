package int64z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/int64z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := int64z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, int64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := int64z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = int64z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, int64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := int64z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = int64z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, int64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, int64(0), int64z.Val(nil))
	require.Equal(t, int64(0), int64z.Val(int64z.Ptr(0)))
	require.Equal(t, int64(1), int64z.Val(int64z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, int64(1), int64z.ValDef(nil, 1))
	require.Equal(t, int64(0), int64z.ValDef(int64z.Ptr(0), 1))
	require.Equal(t, int64(1), int64z.ValDef(int64z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := int64z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, int64(10), v)
	_, err = int64z.Parse("")
	require.Error(t, err)
	_, err = int64z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, int64(10), int64z.MustParse("10")) })
	require.Panics(t, func() { int64z.MustParse("") })
	require.Panics(t, func() { int64z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := int64z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, int64z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[int64]struct{}{}, int64z.SliceToMap(nil))
	require.Equal(t, map[int64]struct{}{}, int64z.SliceToMap([]int64{}))
	require.Equal(t, map[int64]struct{}{1: {}}, int64z.SliceToMap([]int64{1}))
	require.Equal(t, map[int64]struct{}{1: {}, 2: {}}, int64z.SliceToMap([]int64{1, 2}))
	require.Equal(t, map[int64]struct{}{1: {}, 2: {}}, int64z.SliceToMap([]int64{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []int64{}, int64z.MapToSlice(nil))
	require.Equal(t, []int64{}, int64z.MapToSlice(map[int64]struct{}{}))
	require.Equal(t, []int64{1}, int64z.MapToSlice(map[int64]struct{}{1: {}}))
	require.Equal(t, map[int64]struct{}{1: {}, 2: {}}, int64z.SliceToMap(int64z.MapToSlice(map[int64]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, int64(0), int64z.SafeIndex(nil, 0))
	require.Equal(t, int64(0), int64z.SafeIndex(nil, 1))
	require.Equal(t, int64(0), int64z.SafeIndex(nil, -1))
	require.Equal(t, int64(0), int64z.SafeIndex([]int64{}, 0))
	require.Equal(t, int64(0), int64z.SafeIndex([]int64{}, 1))
	require.Equal(t, int64(0), int64z.SafeIndex([]int64{}, -1))
	require.Equal(t, int64(1), int64z.SafeIndex([]int64{1}, 0))
	require.Equal(t, int64(0), int64z.SafeIndex([]int64{1}, 1))
	require.Equal(t, int64(0), int64z.SafeIndex([]int64{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, int64(100), int64z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, int64(100), int64z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, int64(100), int64z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, int64(100), int64z.SafeIndexDef([]int64{}, 0, 100))
	require.Equal(t, int64(100), int64z.SafeIndexDef([]int64{}, 1, 100))
	require.Equal(t, int64(100), int64z.SafeIndexDef([]int64{}, -1, 100))
	require.Equal(t, int64(1), int64z.SafeIndexDef([]int64{1}, 0, 100))
	require.Equal(t, int64(100), int64z.SafeIndexDef([]int64{1}, 1, 100))
	require.Equal(t, int64(100), int64z.SafeIndexDef([]int64{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, int64z.SafeIndexPtr(nil, 0))
	require.Nil(t, int64z.SafeIndexPtr(nil, 1))
	require.Nil(t, int64z.SafeIndexPtr(nil, -1))
	require.Nil(t, int64z.SafeIndexPtr([]int64{}, 0))
	require.Nil(t, int64z.SafeIndexPtr([]int64{}, 1))
	require.Nil(t, int64z.SafeIndexPtr([]int64{}, -1))
	require.Equal(t, int64z.Ptr(1), int64z.SafeIndexPtr([]int64{1}, 0))
	require.Nil(t, int64z.SafeIndexPtr([]int64{1}, 1))
	require.Nil(t, int64z.SafeIndexPtr([]int64{1}, -1))
}
