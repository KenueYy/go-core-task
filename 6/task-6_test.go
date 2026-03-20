package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRangeRandomGenerator(t *testing.T) {
	stopCh := make(chan bool)
	maxValue := 20
	ch := RandomGenerator(stopCh, maxValue)
	nmap := make(map[int]struct{})

	for i := 0; i < 1000; i++ {
		n := <-ch

		nmap[n] = struct{}{}
		require.GreaterOrEqual(t, n, 0)
		require.Less(t, n, maxValue)
	}

	stopCh <- true

	_, flag := <-ch

	require.GreaterOrEqual(t, len(nmap), int(float64(maxValue)*0.9))
	require.False(t, flag)
}
