package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionCheck_HasIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, got := IntersectionCheck(a, b)

	require.True(t, ok)

	want := []int{64, 3}
	require.Equal(t, want, got)
}

func TestIntersectionCheck_NoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	ok, got := IntersectionCheck(a, b)

	require.False(t, ok)
	require.Empty(t, got)
}

func TestIntersectionCheck_EmptySlice(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}

	ok, got := IntersectionCheck(a, b)

	require.False(t, ok)
	require.Empty(t, got)
}
