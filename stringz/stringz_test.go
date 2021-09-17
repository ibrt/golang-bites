package stringz_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-bites/stringz"
)

func TestPtr(t *testing.T) {
	p := stringz.Ptr("")
	require.NotNil(t, p)
	require.Equal(t, "", *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := stringz.PtrZeroToNil("")
	require.Nil(t, p)
	p = stringz.PtrZeroToNil("a")
	require.NotNil(t, p)
	require.Equal(t, "a", *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := stringz.PtrDefToNil("a", "a")
	require.Nil(t, p)
	p = stringz.PtrDefToNil("a", "b")
	require.NotNil(t, p)
	require.Equal(t, "a", *p)
}

func TestVal(t *testing.T) {
	require.Equal(t, "", stringz.Val(nil))
	require.Equal(t, "", stringz.Val(stringz.Ptr("")))
	require.Equal(t, "a", stringz.Val(stringz.Ptr("a")))
}

func TestValDef(t *testing.T) {
	require.Equal(t, "b", stringz.ValDef(nil, "b"))
	require.Equal(t, "", stringz.ValDef(stringz.Ptr(""), "b"))
	require.Equal(t, "a", stringz.ValDef(stringz.Ptr("a"), "b"))
}

func TestSlice(t *testing.T) {
	s := stringz.Slice{"2", "0", "3", "1", "3"}
	require.Equal(t, 5, s.Len())
	require.True(t, s.Less(1, 0))
	require.False(t, s.Less(2, 4))
	require.False(t, s.Less(0, 1))
	s.Swap(0, 1)
	require.Equal(t, stringz.Slice{"0", "2", "3", "1", "3"}, s)
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[string]struct{}{}, stringz.SliceToMap(nil))
	require.Equal(t, map[string]struct{}{}, stringz.SliceToMap([]string{}))
	require.Equal(t, map[string]struct{}{"1": {}}, stringz.SliceToMap([]string{"1"}))
	require.Equal(t, map[string]struct{}{"1": {}, "2": {}}, stringz.SliceToMap([]string{"1", "2"}))
	require.Equal(t, map[string]struct{}{"1": {}, "2": {}}, stringz.SliceToMap([]string{"1", "1", "2", "2"}))
}

func TestMapToSlice(t *testing.T) {
	require.Equal(t, []string{}, stringz.MapToSlice(nil))
	require.Equal(t, []string{}, stringz.MapToSlice(map[string]struct{}{}))
	require.Equal(t, []string{"1"}, stringz.MapToSlice(map[string]struct{}{"1": {}}))
	require.Equal(t, map[string]struct{}{"1": {}, "2": {}}, stringz.SliceToMap(stringz.MapToSlice(map[string]struct{}{"1": {}, "2": {}})))
}

func TestSafeIndex(t *testing.T) {
	require.Equal(t, "", stringz.SafeIndex(nil, 0))
	require.Equal(t, "", stringz.SafeIndex(nil, 1))
	require.Equal(t, "", stringz.SafeIndex(nil, -1))
	require.Equal(t, "", stringz.SafeIndex([]string{}, 0))
	require.Equal(t, "", stringz.SafeIndex([]string{}, 1))
	require.Equal(t, "", stringz.SafeIndex([]string{}, -1))
	require.Equal(t, "a", stringz.SafeIndex([]string{"a"}, 0))
	require.Equal(t, "", stringz.SafeIndex([]string{"a"}, 1))
	require.Equal(t, "", stringz.SafeIndex([]string{"a"}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.Equal(t, "b", stringz.SafeIndexDef(nil, 0, "b"))
	require.Equal(t, "b", stringz.SafeIndexDef(nil, 1, "b"))
	require.Equal(t, "b", stringz.SafeIndexDef(nil, -1, "b"))
	require.Equal(t, "b", stringz.SafeIndexDef([]string{}, 0, "b"))
	require.Equal(t, "b", stringz.SafeIndexDef([]string{}, 1, "b"))
	require.Equal(t, "b", stringz.SafeIndexDef([]string{}, -1, "b"))
	require.Equal(t, "a", stringz.SafeIndexDef([]string{"a"}, 0, "b"))
	require.Equal(t, "b", stringz.SafeIndexDef([]string{"a"}, 1, "b"))
	require.Equal(t, "b", stringz.SafeIndexDef([]string{"a"}, -1, "b"))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, stringz.SafeIndexPtr(nil, 0))
	require.Nil(t, stringz.SafeIndexPtr(nil, 1))
	require.Nil(t, stringz.SafeIndexPtr(nil, -1))
	require.Nil(t, stringz.SafeIndexPtr([]string{}, 0))
	require.Nil(t, stringz.SafeIndexPtr([]string{}, 1))
	require.Nil(t, stringz.SafeIndexPtr([]string{}, -1))
	require.Equal(t, stringz.Ptr("a"), stringz.SafeIndexPtr([]string{"a"}, 0))
	require.Nil(t, stringz.SafeIndexPtr([]string{"a"}, 1))
	require.Nil(t, stringz.SafeIndexPtr([]string{"a"}, -1))
}
