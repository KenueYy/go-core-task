package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewCustomWaitGroup_SizeClamp(t *testing.T) {
	wg := NewCustomWaitGroup(0)
	require.Equal(t, 1, cap(wg.sem))

	wg = NewCustomWaitGroup(-5)
	require.Equal(t, 1, cap(wg.sem))

	wg = NewCustomWaitGroup(3)
	require.Equal(t, 3, cap(wg.sem))
}

func TestCustomWaitGroup_AddDoneBalance(t *testing.T) {
	wg := NewCustomWaitGroup(3)

	wg.Add()
	wg.Add()
	wg.Add()

	wg.Done()
	wg.Done()
	wg.Done()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Add/done misscount")
	}
}

func TestCustomWaitGroup_Concurrent(t *testing.T) {
	const (
		size       = 5
		goroutines = 20
	)

	wg := NewCustomWaitGroup(size)
	var internal sync.WaitGroup
	internal.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer internal.Done()

			wg.Add()
			time.Sleep(5 * time.Millisecond)
			wg.Done()
		}()
	}

	internal.Wait()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("not all goroutines finished")
	}
}
