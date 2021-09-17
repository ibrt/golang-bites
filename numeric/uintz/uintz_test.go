package uintz_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/uintz"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uintz.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uintz.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uintz.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uintz.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uintz.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint(0), uintz.Val(nil))
	require.Equal(t, uint(0), uintz.Val(uintz.Ptr(0)))
	require.Equal(t, uint(1), uintz.Val(uintz.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint(1), uintz.ValDef(nil, 1))
	require.Equal(t, uint(0), uintz.ValDef(uintz.Ptr(0), 1))
	require.Equal(t, uint(1), uintz.ValDef(uintz.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uintz.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint(10), v)
	_, err = uintz.Parse("")
	require.Error(t, err)
	_, err = uintz.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, uint(10), uintz.MustParse("10")) })
	require.Panics(t, func() { uintz.MustParse("") })
	require.Panics(t, func() { uintz.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := uintz.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, uintz.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint]struct{}{}, uintz.SliceToMap(nil))
	require.Equal(t, map[uint]struct{}{}, uintz.SliceToMap([]uint{}))
	require.Equal(t, map[uint]struct{}{1: {}}, uintz.SliceToMap([]uint{1}))
	require.Equal(t, map[uint]struct{}{1: {}, 2: {}}, uintz.SliceToMap([]uint{1, 2}))
	require.Equal(t, map[uint]struct{}{1: {}, 2: {}}, uintz.SliceToMap([]uint{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint{}, uintz.MapToSlice(nil))
	require.Equal(t, []uint{}, uintz.MapToSlice(map[uint]struct{}{}))
	require.Equal(t, []uint{1}, uintz.MapToSlice(map[uint]struct{}{1: {}}))
	require.Equal(t, map[uint]struct{}{1: {}, 2: {}}, uintz.SliceToMap(uintz.MapToSlice(map[uint]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, uint(0), uintz.SafeIndex(nil, 0))
	require.Equal(t, uint(0), uintz.SafeIndex(nil, 1))
	require.Equal(t, uint(0), uintz.SafeIndex(nil, -1))
	require.Equal(t, uint(0), uintz.SafeIndex([]uint{}, 0))
	require.Equal(t, uint(0), uintz.SafeIndex([]uint{}, 1))
	require.Equal(t, uint(0), uintz.SafeIndex([]uint{}, -1))
	require.Equal(t, uint(1), uintz.SafeIndex([]uint{1}, 0))
	require.Equal(t, uint(0), uintz.SafeIndex([]uint{1}, 1))
	require.Equal(t, uint(0), uintz.SafeIndex([]uint{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, uint(100), uintz.SafeIndexDef(nil, 0, 100))
	require.Equal(t, uint(100), uintz.SafeIndexDef(nil, 1, 100))
	require.Equal(t, uint(100), uintz.SafeIndexDef(nil, -1, 100))
	require.Equal(t, uint(100), uintz.SafeIndexDef([]uint{}, 0, 100))
	require.Equal(t, uint(100), uintz.SafeIndexDef([]uint{}, 1, 100))
	require.Equal(t, uint(100), uintz.SafeIndexDef([]uint{}, -1, 100))
	require.Equal(t, uint(1), uintz.SafeIndexDef([]uint{1}, 0, 100))
	require.Equal(t, uint(100), uintz.SafeIndexDef([]uint{1}, 1, 100))
	require.Equal(t, uint(100), uintz.SafeIndexDef([]uint{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, uintz.SafeIndexPtr(nil, 0))
	require.Nil(t, uintz.SafeIndexPtr(nil, 1))
	require.Nil(t, uintz.SafeIndexPtr(nil, -1))
	require.Nil(t, uintz.SafeIndexPtr([]uint{}, 0))
	require.Nil(t, uintz.SafeIndexPtr([]uint{}, 1))
	require.Nil(t, uintz.SafeIndexPtr([]uint{}, -1))
	require.Equal(t, uintz.Ptr(1), uintz.SafeIndexPtr([]uint{1}, 0))
	require.Nil(t, uintz.SafeIndexPtr([]uint{1}, 1))
	require.Nil(t, uintz.SafeIndexPtr([]uint{1}, -1))
}
