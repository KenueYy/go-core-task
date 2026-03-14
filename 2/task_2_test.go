package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOriginalSlice_LengthAndRandom(t *testing.T) {
	s1 := originalSlice(10)
	s2 := originalSlice(10)

	require.Len(t, s1, 10)
	require.Len(t, s2, 10)

	require.NotEqual(t, s1, s2)
}

func TestSliceExample_EvenNumbers(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "mixed",
			in:   []int{1, 2, 3, 4, 5, 6},
			want: []int{2, 4, 6},
		},
		{
			name: "all odd",
			in:   []int{1, 3, 5},
			want: []int{},
		},
		{
			name: "all even",
			in:   []int{2, 4, 6},
			want: []int{2, 4, 6},
		},
		{
			name: "empty",
			in:   []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.in)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestAddElement(t *testing.T) {
	tests := []struct {
		name   string
		in     []int
		number int
		want   []int
	}{
		{
			name:   "add to non-empty",
			in:     []int{1, 2, 3},
			number: 10,
			want:   []int{1, 2, 3, 10},
		},
		{
			name:   "add to empty",
			in:     []int{},
			number: 5,
			want:   []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElement(tt.in, tt.number)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCopySlice_IndependentCopy(t *testing.T) {
	src := []int{1, 2, 3, 4}
	cp := copySlice(src)

	require.Equal(t, src, cp)
	require.NotSame(t, &src[0], &cp[0])

	src[0] = 99
	require.Equal(t, []int{99, 2, 3, 4}, src)
	require.Equal(t, []int{1, 2, 3, 4}, cp)
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		pos  int
		want []int
	}{
		{
			name: "remove middle",
			in:   []int{1, 2, 3, 4, 5},
			pos:  2,
			want: []int{1, 2, 4, 5},
		},
		{
			name: "remove first",
			in:   []int{1, 2, 3},
			pos:  0,
			want: []int{2, 3},
		},
		{
			name: "remove last",
			in:   []int{1, 2, 3},
			pos:  2,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.in, tt.pos)
			require.Equal(t, tt.want, got)
		})
	}
}
