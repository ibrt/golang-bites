package uint8z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/uint8z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint8z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint8(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint8z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint8z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint8(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint8z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint8z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint8(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint8(0), uint8z.Val(nil))
	require.Equal(t, uint8(0), uint8z.Val(uint8z.Ptr(0)))
	require.Equal(t, uint8(1), uint8z.Val(uint8z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint8(1), uint8z.ValDef(nil, 1))
	require.Equal(t, uint8(0), uint8z.ValDef(uint8z.Ptr(0), 1))
	require.Equal(t, uint8(1), uint8z.ValDef(uint8z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint8z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint8(10), v)
	_, err = uint8z.Parse("")
	require.Error(t, err)
	_, err = uint8z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, uint8(10), uint8z.MustParse("10")) })
	require.Panics(t, func() { uint8z.MustParse("") })
	require.Panics(t, func() { uint8z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := uint8z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, uint8z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint8]struct{}{}, uint8z.SliceToMap(nil))
	require.Equal(t, map[uint8]struct{}{}, uint8z.SliceToMap([]uint8{}))
	require.Equal(t, map[uint8]struct{}{1: {}}, uint8z.SliceToMap([]uint8{1}))
	require.Equal(t, map[uint8]struct{}{1: {}, 2: {}}, uint8z.SliceToMap([]uint8{1, 2}))
	require.Equal(t, map[uint8]struct{}{1: {}, 2: {}}, uint8z.SliceToMap([]uint8{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint8{}, uint8z.MapToSlice(nil))
	require.Equal(t, []uint8{}, uint8z.MapToSlice(map[uint8]struct{}{}))
	require.Equal(t, []uint8{1}, uint8z.MapToSlice(map[uint8]struct{}{1: {}}))
	require.Equal(t, map[uint8]struct{}{1: {}, 2: {}}, uint8z.SliceToMap(uint8z.MapToSlice(map[uint8]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, uint8(0), uint8z.SafeIndex(nil, 0))
	require.Equal(t, uint8(0), uint8z.SafeIndex(nil, 1))
	require.Equal(t, uint8(0), uint8z.SafeIndex(nil, -1))
	require.Equal(t, uint8(0), uint8z.SafeIndex([]uint8{}, 0))
	require.Equal(t, uint8(0), uint8z.SafeIndex([]uint8{}, 1))
	require.Equal(t, uint8(0), uint8z.SafeIndex([]uint8{}, -1))
	require.Equal(t, uint8(1), uint8z.SafeIndex([]uint8{1}, 0))
	require.Equal(t, uint8(0), uint8z.SafeIndex([]uint8{1}, 1))
	require.Equal(t, uint8(0), uint8z.SafeIndex([]uint8{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, uint8(100), uint8z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, uint8(100), uint8z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, uint8(100), uint8z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, uint8(100), uint8z.SafeIndexDef([]uint8{}, 0, 100))
	require.Equal(t, uint8(100), uint8z.SafeIndexDef([]uint8{}, 1, 100))
	require.Equal(t, uint8(100), uint8z.SafeIndexDef([]uint8{}, -1, 100))
	require.Equal(t, uint8(1), uint8z.SafeIndexDef([]uint8{1}, 0, 100))
	require.Equal(t, uint8(100), uint8z.SafeIndexDef([]uint8{1}, 1, 100))
	require.Equal(t, uint8(100), uint8z.SafeIndexDef([]uint8{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, uint8z.SafeIndexPtr(nil, 0))
	require.Nil(t, uint8z.SafeIndexPtr(nil, 1))
	require.Nil(t, uint8z.SafeIndexPtr(nil, -1))
	require.Nil(t, uint8z.SafeIndexPtr([]uint8{}, 0))
	require.Nil(t, uint8z.SafeIndexPtr([]uint8{}, 1))
	require.Nil(t, uint8z.SafeIndexPtr([]uint8{}, -1))
	require.Equal(t, uint8z.Ptr(1), uint8z.SafeIndexPtr([]uint8{1}, 0))
	require.Nil(t, uint8z.SafeIndexPtr([]uint8{1}, 1))
	require.Nil(t, uint8z.SafeIndexPtr([]uint8{1}, -1))
}
