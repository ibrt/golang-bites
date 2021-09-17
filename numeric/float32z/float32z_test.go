package float32z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/float32z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := float32z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, float32(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := float32z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = float32z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, float32(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := float32z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = float32z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, float32(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, float32(0), float32z.Val(nil))
	require.Equal(t, float32(0), float32z.Val(float32z.Ptr(0)))
	require.Equal(t, float32(1), float32z.Val(float32z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, float32(1), float32z.ValDef(nil, 1))
	require.Equal(t, float32(0), float32z.ValDef(float32z.Ptr(0), 1))
	require.Equal(t, float32(1), float32z.ValDef(float32z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := float32z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, float32(10), v)
	_, err = float32z.Parse("")
	require.Error(t, err)
	_, err = float32z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, float32(10), float32z.MustParse("10")) })
	require.Panics(t, func() { float32z.MustParse("") })
	require.Panics(t, func() { float32z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := float32z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, float32z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[float32]struct{}{}, float32z.SliceToMap(nil))
	require.Equal(t, map[float32]struct{}{}, float32z.SliceToMap([]float32{}))
	require.Equal(t, map[float32]struct{}{1: {}}, float32z.SliceToMap([]float32{1}))
	require.Equal(t, map[float32]struct{}{1: {}, 2: {}}, float32z.SliceToMap([]float32{1, 2}))
	require.Equal(t, map[float32]struct{}{1: {}, 2: {}}, float32z.SliceToMap([]float32{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []float32{}, float32z.MapToSlice(nil))
	require.Equal(t, []float32{}, float32z.MapToSlice(map[float32]struct{}{}))
	require.Equal(t, []float32{1}, float32z.MapToSlice(map[float32]struct{}{1: {}}))
	require.Equal(t, map[float32]struct{}{1: {}, 2: {}}, float32z.SliceToMap(float32z.MapToSlice(map[float32]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, float32(0), float32z.SafeIndex(nil, 0))
	require.Equal(t, float32(0), float32z.SafeIndex(nil, 1))
	require.Equal(t, float32(0), float32z.SafeIndex(nil, -1))
	require.Equal(t, float32(0), float32z.SafeIndex([]float32{}, 0))
	require.Equal(t, float32(0), float32z.SafeIndex([]float32{}, 1))
	require.Equal(t, float32(0), float32z.SafeIndex([]float32{}, -1))
	require.Equal(t, float32(1), float32z.SafeIndex([]float32{1}, 0))
	require.Equal(t, float32(0), float32z.SafeIndex([]float32{1}, 1))
	require.Equal(t, float32(0), float32z.SafeIndex([]float32{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, float32(100), float32z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, float32(100), float32z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, float32(100), float32z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, float32(100), float32z.SafeIndexDef([]float32{}, 0, 100))
	require.Equal(t, float32(100), float32z.SafeIndexDef([]float32{}, 1, 100))
	require.Equal(t, float32(100), float32z.SafeIndexDef([]float32{}, -1, 100))
	require.Equal(t, float32(1), float32z.SafeIndexDef([]float32{1}, 0, 100))
	require.Equal(t, float32(100), float32z.SafeIndexDef([]float32{1}, 1, 100))
	require.Equal(t, float32(100), float32z.SafeIndexDef([]float32{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, float32z.SafeIndexPtr(nil, 0))
	require.Nil(t, float32z.SafeIndexPtr(nil, 1))
	require.Nil(t, float32z.SafeIndexPtr(nil, -1))
	require.Nil(t, float32z.SafeIndexPtr([]float32{}, 0))
	require.Nil(t, float32z.SafeIndexPtr([]float32{}, 1))
	require.Nil(t, float32z.SafeIndexPtr([]float32{}, -1))
	require.Equal(t, float32z.Ptr(1), float32z.SafeIndexPtr([]float32{1}, 0))
	require.Nil(t, float32z.SafeIndexPtr([]float32{1}, 1))
	require.Nil(t, float32z.SafeIndexPtr([]float32{1}, -1))
}
