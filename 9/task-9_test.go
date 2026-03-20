package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Cube(t *testing.T) {
	intChan := make(chan uint)

	go func() {
		defer close(intChan)
		for i := 0; i < 10; i++ {
			intChan <- uint(i)
		}
	}()

	cubed := Cube(intChan)

	expected := map[uint]uint{
		0: 0, 1: 1, 2: 8, 3: 27, 4: 64,
		5: 125, 6: 216, 7: 343, 8: 512, 9: 729,
	}

	got := make(map[uint]uint)
	var i uint = 0
	for num := range cubed {
		got[i] = num
		require.Less(t, num, uint(1000))
		i++
	}

	require.Equal(t, expected, got)
}

func Test_Convert(t *testing.T) {
	intChan := make(chan uint)
	go func() {
		defer close(intChan)
		for i := 0; i < 10; i++ {
			intChan <- uint(i)
		}
	}()

	out := FloatConverter(intChan)

	var i uint
	for fl := range out {
		require.IsType(t, float64(0), fl)
		require.Equal(t, float64(i), fl)
		i++
	}
}
