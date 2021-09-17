package uint64z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/uint64z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := uint64z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, uint64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := uint64z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = uint64z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, uint64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := uint64z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = uint64z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, uint64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, uint64(0), uint64z.Val(nil))
	require.Equal(t, uint64(0), uint64z.Val(uint64z.Ptr(0)))
	require.Equal(t, uint64(1), uint64z.Val(uint64z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, uint64(1), uint64z.ValDef(nil, 1))
	require.Equal(t, uint64(0), uint64z.ValDef(uint64z.Ptr(0), 1))
	require.Equal(t, uint64(1), uint64z.ValDef(uint64z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := uint64z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, uint64(10), v)
	_, err = uint64z.Parse("")
	require.Error(t, err)
	_, err = uint64z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, uint64(10), uint64z.MustParse("10")) })
	require.Panics(t, func() { uint64z.MustParse("") })
	require.Panics(t, func() { uint64z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := uint64z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, uint64z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[uint64]struct{}{}, uint64z.SliceToMap(nil))
	require.Equal(t, map[uint64]struct{}{}, uint64z.SliceToMap([]uint64{}))
	require.Equal(t, map[uint64]struct{}{1: {}}, uint64z.SliceToMap([]uint64{1}))
	require.Equal(t, map[uint64]struct{}{1: {}, 2: {}}, uint64z.SliceToMap([]uint64{1, 2}))
	require.Equal(t, map[uint64]struct{}{1: {}, 2: {}}, uint64z.SliceToMap([]uint64{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []uint64{}, uint64z.MapToSlice(nil))
	require.Equal(t, []uint64{}, uint64z.MapToSlice(map[uint64]struct{}{}))
	require.Equal(t, []uint64{1}, uint64z.MapToSlice(map[uint64]struct{}{1: {}}))
	require.Equal(t, map[uint64]struct{}{1: {}, 2: {}}, uint64z.SliceToMap(uint64z.MapToSlice(map[uint64]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, uint64(0), uint64z.SafeIndex(nil, 0))
	require.Equal(t, uint64(0), uint64z.SafeIndex(nil, 1))
	require.Equal(t, uint64(0), uint64z.SafeIndex(nil, -1))
	require.Equal(t, uint64(0), uint64z.SafeIndex([]uint64{}, 0))
	require.Equal(t, uint64(0), uint64z.SafeIndex([]uint64{}, 1))
	require.Equal(t, uint64(0), uint64z.SafeIndex([]uint64{}, -1))
	require.Equal(t, uint64(1), uint64z.SafeIndex([]uint64{1}, 0))
	require.Equal(t, uint64(0), uint64z.SafeIndex([]uint64{1}, 1))
	require.Equal(t, uint64(0), uint64z.SafeIndex([]uint64{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, uint64(100), uint64z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, uint64(100), uint64z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, uint64(100), uint64z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, uint64(100), uint64z.SafeIndexDef([]uint64{}, 0, 100))
	require.Equal(t, uint64(100), uint64z.SafeIndexDef([]uint64{}, 1, 100))
	require.Equal(t, uint64(100), uint64z.SafeIndexDef([]uint64{}, -1, 100))
	require.Equal(t, uint64(1), uint64z.SafeIndexDef([]uint64{1}, 0, 100))
	require.Equal(t, uint64(100), uint64z.SafeIndexDef([]uint64{1}, 1, 100))
	require.Equal(t, uint64(100), uint64z.SafeIndexDef([]uint64{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, uint64z.SafeIndexPtr(nil, 0))
	require.Nil(t, uint64z.SafeIndexPtr(nil, 1))
	require.Nil(t, uint64z.SafeIndexPtr(nil, -1))
	require.Nil(t, uint64z.SafeIndexPtr([]uint64{}, 0))
	require.Nil(t, uint64z.SafeIndexPtr([]uint64{}, 1))
	require.Nil(t, uint64z.SafeIndexPtr([]uint64{}, -1))
	require.Equal(t, uint64z.Ptr(1), uint64z.SafeIndexPtr([]uint64{1}, 0))
	require.Nil(t, uint64z.SafeIndexPtr([]uint64{1}, 1))
	require.Nil(t, uint64z.SafeIndexPtr([]uint64{1}, -1))
}
