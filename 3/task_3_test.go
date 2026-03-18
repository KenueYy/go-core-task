package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStringIntMap_Init(t *testing.T) {
	m := NewStringIntMap()
	require.NotNil(t, m)
	require.NotNil(t, m.data)
	require.Len(t, m.data, 0)
}

func TestStringIntMap_AddAndGet(t *testing.T) {
	m := NewStringIntMap()

	m.Add("a", 1)
	m.Add("b", 2)

	v, ok := m.Get("a")
	require.True(t, ok)
	require.Equal(t, 1, v)

	v, ok = m.Get("b")
	require.True(t, ok)
	require.Equal(t, 2, v)

	v, ok = m.Get("c")
	require.False(t, ok)
	require.Equal(t, 0, v)
}

func TestStringIntMap_Exists(t *testing.T) {
	m := NewStringIntMap()

	require.False(t, m.Exists("x"))

	m.Add("x", 10)
	require.True(t, m.Exists("x"))

	m.Remove("x")
	require.False(t, m.Exists("x"))
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("k1", 100)
	m.Add("k2", 200)

	require.True(t, m.Exists("k1"))
	require.True(t, m.Exists("k2"))

	m.Remove("k1")
	require.False(t, m.Exists("k1"))
	require.True(t, m.Exists("k2"))

	m.Remove("no_such_key")
}

func TestStringIntMap_Copy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	cp := m.Copy()

	require.Equal(t, map[string]int{"a": 1, "b": 2}, cp)

	cp["a"] = 999
	v, _ := m.Get("a")
	require.Equal(t, 1, v)

	m.Add("c", 3)
	require.NotContains(t, cp, "c")
}
