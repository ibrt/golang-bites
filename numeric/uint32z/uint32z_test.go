package uint32z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/uint32z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint32z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint32z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint32z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint32z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint32z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint32(0), uint32z.Val(nil))
	require.Equal(t, uint32(0), uint32z.Val(uint32z.Ptr(0)))
	require.Equal(t, uint32(1), uint32z.Val(uint32z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint32(1), uint32z.ValDef(nil, 1))
	require.Equal(t, uint32(0), uint32z.ValDef(uint32z.Ptr(0), 1))
	require.Equal(t, uint32(1), uint32z.ValDef(uint32z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint32z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint32(10), v)
	_, err = uint32z.Parse("")
	require.Error(t, err)
	_, err = uint32z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, uint32(10), uint32z.MustParse("10")) })
	require.Panics(t, func() { uint32z.MustParse("") })
	require.Panics(t, func() { uint32z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := uint32z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, uint32z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint32]struct{}{}, uint32z.SliceToMap(nil))
	require.Equal(t, map[uint32]struct{}{}, uint32z.SliceToMap([]uint32{}))
	require.Equal(t, map[uint32]struct{}{1: {}}, uint32z.SliceToMap([]uint32{1}))
	require.Equal(t, map[uint32]struct{}{1: {}, 2: {}}, uint32z.SliceToMap([]uint32{1, 2}))
	require.Equal(t, map[uint32]struct{}{1: {}, 2: {}}, uint32z.SliceToMap([]uint32{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint32{}, uint32z.MapToSlice(nil))
	require.Equal(t, []uint32{}, uint32z.MapToSlice(map[uint32]struct{}{}))
	require.Equal(t, []uint32{1}, uint32z.MapToSlice(map[uint32]struct{}{1: {}}))
	require.Equal(t, map[uint32]struct{}{1: {}, 2: {}}, uint32z.SliceToMap(uint32z.MapToSlice(map[uint32]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, uint32(0), uint32z.SafeIndex(nil, 0))
	require.Equal(t, uint32(0), uint32z.SafeIndex(nil, 1))
	require.Equal(t, uint32(0), uint32z.SafeIndex(nil, -1))
	require.Equal(t, uint32(0), uint32z.SafeIndex([]uint32{}, 0))
	require.Equal(t, uint32(0), uint32z.SafeIndex([]uint32{}, 1))
	require.Equal(t, uint32(0), uint32z.SafeIndex([]uint32{}, -1))
	require.Equal(t, uint32(1), uint32z.SafeIndex([]uint32{1}, 0))
	require.Equal(t, uint32(0), uint32z.SafeIndex([]uint32{1}, 1))
	require.Equal(t, uint32(0), uint32z.SafeIndex([]uint32{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, uint32(100), uint32z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, uint32(100), uint32z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, uint32(100), uint32z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, uint32(100), uint32z.SafeIndexDef([]uint32{}, 0, 100))
	require.Equal(t, uint32(100), uint32z.SafeIndexDef([]uint32{}, 1, 100))
	require.Equal(t, uint32(100), uint32z.SafeIndexDef([]uint32{}, -1, 100))
	require.Equal(t, uint32(1), uint32z.SafeIndexDef([]uint32{1}, 0, 100))
	require.Equal(t, uint32(100), uint32z.SafeIndexDef([]uint32{1}, 1, 100))
	require.Equal(t, uint32(100), uint32z.SafeIndexDef([]uint32{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, uint32z.SafeIndexPtr(nil, 0))
	require.Nil(t, uint32z.SafeIndexPtr(nil, 1))
	require.Nil(t, uint32z.SafeIndexPtr(nil, -1))
	require.Nil(t, uint32z.SafeIndexPtr([]uint32{}, 0))
	require.Nil(t, uint32z.SafeIndexPtr([]uint32{}, 1))
	require.Nil(t, uint32z.SafeIndexPtr([]uint32{}, -1))
	require.Equal(t, uint32z.Ptr(1), uint32z.SafeIndexPtr([]uint32{1}, 0))
	require.Nil(t, uint32z.SafeIndexPtr([]uint32{1}, 1))
	require.Nil(t, uint32z.SafeIndexPtr([]uint32{1}, -1))
}
