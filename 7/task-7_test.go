package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeChannel_AllValuesPresent(t *testing.T) {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 4)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	ch2 <- 10
	ch2 <- 20
	close(ch2)

	ch3 <- 100
	ch3 <- 200
	ch3 <- 300
	ch3 <- 400
	close(ch3)

	out := MergeChannel(ch1, ch2, ch3)

	var result []int
	for v := range out {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 10, 20, 100, 200, 300, 400}

	require.Len(t, result, len(expected))

	toMultiset := func(vals []int) map[int]int {
		m := make(map[int]int)
		for _, v := range vals {
			m[v]++
		}
		return m
	}

	require.Equal(t, toMultiset(expected), toMultiset(result))
}
