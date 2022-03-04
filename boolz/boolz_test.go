package boolz_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-bites/boolz"
)

func TestPtr(t *testing.T) {
	p := boolz.Ptr(false)
	require.NotNil(t, p)
	require.False(t, *p)

	p = boolz.Ptr(true)
	require.NotNil(t, p)
	require.True(t, *p)
}

func TestPtrZeroToNil(t *testing.T) {
	p := boolz.PtrZeroToNil(false)
	require.Nil(t, p)
	p = boolz.PtrZeroToNil(true)
	require.NotNil(t, p)
	require.True(t, *p)
}

func TestPtrDefToNil(t *testing.T) {
	p := boolz.PtrDefToNil(true, true)
	require.Nil(t, p)
	p = boolz.PtrDefToNil(true, false)
	require.NotNil(t, p)
	require.True(t, *p)
	p = boolz.PtrDefToNil(false, true)
	require.NotNil(t, p)
	require.False(t, *p)
}

func TestVal(t *testing.T) {
	require.False(t, boolz.Val(nil))
	require.False(t, boolz.Val(boolz.Ptr(false)))
	require.True(t, boolz.Val(boolz.Ptr(true)))
}

func TestValDef(t *testing.T) {
	require.True(t, boolz.ValDef(nil, true))
	require.False(t, boolz.ValDef(boolz.Ptr(false), true))
	require.True(t, boolz.ValDef(boolz.Ptr(true), true))
}

func TestSafeIndex(t *testing.T) {
	require.False(t, boolz.SafeIndex(nil, 0))
	require.False(t, boolz.SafeIndex(nil, 1))
	require.False(t, boolz.SafeIndex(nil, -1))
	require.False(t, boolz.SafeIndex([]bool{}, 0))
	require.False(t, boolz.SafeIndex([]bool{}, 1))
	require.False(t, boolz.SafeIndex([]bool{}, -1))
	require.True(t, boolz.SafeIndex([]bool{true}, 0))
	require.False(t, boolz.SafeIndex([]bool{true}, 1))
	require.False(t, boolz.SafeIndex([]bool{true}, -1))
}

func TestSafeIndexDef(t *testing.T) {
	require.True(t, boolz.SafeIndexDef(nil, 0, true))
	require.True(t, boolz.SafeIndexDef(nil, 1, true))
	require.True(t, boolz.SafeIndexDef(nil, -1, true))
	require.True(t, boolz.SafeIndexDef([]bool{}, 0, true))
	require.True(t, boolz.SafeIndexDef([]bool{}, 1, true))
	require.True(t, boolz.SafeIndexDef([]bool{}, -1, true))
	require.False(t, boolz.SafeIndexDef([]bool{false}, 0, true))
	require.True(t, boolz.SafeIndexDef([]bool{false}, 1, true))
	require.True(t, boolz.SafeIndexDef([]bool{false}, -1, true))
}

func TestSafeIndexPtr(t *testing.T) {
	require.Nil(t, boolz.SafeIndexPtr(nil, 0))
	require.Nil(t, boolz.SafeIndexPtr(nil, 1))
	require.Nil(t, boolz.SafeIndexPtr(nil, -1))
	require.Nil(t, boolz.SafeIndexPtr([]bool{}, 0))
	require.Nil(t, boolz.SafeIndexPtr([]bool{}, 1))
	require.Nil(t, boolz.SafeIndexPtr([]bool{}, -1))
	require.Equal(t, boolz.Ptr(true), boolz.SafeIndexPtr([]bool{true}, 0))
	require.Nil(t, boolz.SafeIndexPtr([]bool{true}, 1))
	require.Nil(t, boolz.SafeIndexPtr([]bool{true}, -1))
}
