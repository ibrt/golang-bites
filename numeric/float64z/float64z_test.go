package float64z_test

import (
	"testing"

	"github.com/ibrt/golang-bites/numeric/float64z"

	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	p := float64z.Ptr(0)
	require.NotNil(t, p)
	require.Equal(t, float64(0), *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := float64z.PtrZeroToNil(0)
	require.Nil(t, p)
	p = float64z.PtrZeroToNil(1)
	require.NotNil(t, p)
	require.Equal(t, float64(1), *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := float64z.PtrDefToNil(1, 1)
	require.Nil(t, p)
	p = float64z.PtrDefToNil(1, 0)
	require.NotNil(t, p)
	require.Equal(t, float64(1), *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, float64(0), float64z.Val(nil))
	require.Equal(t, float64(0), float64z.Val(float64z.Ptr(0)))
	require.Equal(t, float64(1), float64z.Val(float64z.Ptr(1)))
}

func TestValDef(t *testing.T) {
	require.Equal(t, float64(1), float64z.ValDef(nil, 1))
	require.Equal(t, float64(0), float64z.ValDef(float64z.Ptr(0), 1))
	require.Equal(t, float64(1), float64z.ValDef(float64z.Ptr(1), 1))
}

func TestParse(t *testing.T) {
	v, err := float64z.Parse("10")
	require.NoError(t, err)
	require.Equal(t, float64(10), v)
	_, err = float64z.Parse("")
	require.Error(t, err)
	_, err = float64z.Parse("A")
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {
	require.NotPanics(t, func() { require.Equal(t, float64(10), float64z.MustParse("10")) })
	require.Panics(t, func() { float64z.MustParse("") })
	require.Panics(t, func() { float64z.MustParse("A") })
}

func TestSlice(t *testing.T) {
	s := float64z.Slice{2, 0, 3, 1, 3}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, float64z.Slice{0, 2, 3, 1, 3}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[float64]struct{}{}, float64z.SliceToMap(nil))
	require.Equal(t, map[float64]struct{}{}, float64z.SliceToMap([]float64{}))
	require.Equal(t, map[float64]struct{}{1: {}}, float64z.SliceToMap([]float64{1}))
	require.Equal(t, map[float64]struct{}{1: {}, 2: {}}, float64z.SliceToMap([]float64{1, 2}))
	require.Equal(t, map[float64]struct{}{1: {}, 2: {}}, float64z.SliceToMap([]float64{1, 1, 2, 2}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []float64{}, float64z.MapToSlice(nil))
	require.Equal(t, []float64{}, float64z.MapToSlice(map[float64]struct{}{}))
	require.Equal(t, []float64{1}, float64z.MapToSlice(map[float64]struct{}{1: {}}))
	require.Equal(t, map[float64]struct{}{1: {}, 2: {}}, float64z.SliceToMap(float64z.MapToSlice(map[float64]struct{}{1: {}, 2: {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, float64(0), float64z.SafeIndex(nil, 0))
	require.Equal(t, float64(0), float64z.SafeIndex(nil, 1))
	require.Equal(t, float64(0), float64z.SafeIndex(nil, -1))
	require.Equal(t, float64(0), float64z.SafeIndex([]float64{}, 0))
	require.Equal(t, float64(0), float64z.SafeIndex([]float64{}, 1))
	require.Equal(t, float64(0), float64z.SafeIndex([]float64{}, -1))
	require.Equal(t, float64(1), float64z.SafeIndex([]float64{1}, 0))
	require.Equal(t, float64(0), float64z.SafeIndex([]float64{1}, 1))
	require.Equal(t, float64(0), float64z.SafeIndex([]float64{1}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, float64(100), float64z.SafeIndexDef(nil, 0, 100))
	require.Equal(t, float64(100), float64z.SafeIndexDef(nil, 1, 100))
	require.Equal(t, float64(100), float64z.SafeIndexDef(nil, -1, 100))
	require.Equal(t, float64(100), float64z.SafeIndexDef([]float64{}, 0, 100))
	require.Equal(t, float64(100), float64z.SafeIndexDef([]float64{}, 1, 100))
	require.Equal(t, float64(100), float64z.SafeIndexDef([]float64{}, -1, 100))
	require.Equal(t, float64(1), float64z.SafeIndexDef([]float64{1}, 0, 100))
	require.Equal(t, float64(100), float64z.SafeIndexDef([]float64{1}, 1, 100))
	require.Equal(t, float64(100), float64z.SafeIndexDef([]float64{1}, -1, 100))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, float64z.SafeIndexPtr(nil, 0))
	require.Nil(t, float64z.SafeIndexPtr(nil, 1))
	require.Nil(t, float64z.SafeIndexPtr(nil, -1))
	require.Nil(t, float64z.SafeIndexPtr([]float64{}, 0))
	require.Nil(t, float64z.SafeIndexPtr([]float64{}, 1))
	require.Nil(t, float64z.SafeIndexPtr([]float64{}, -1))
	require.Equal(t, float64z.Ptr(1), float64z.SafeIndexPtr([]float64{1}, 0))
	require.Nil(t, float64z.SafeIndexPtr([]float64{1}, 1))
	require.Nil(t, float64z.SafeIndexPtr([]float64{1}, -1))
}
