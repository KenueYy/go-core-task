package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	fchan := make(chan uint)
	schan := make(chan float64)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 0; i < 3; i++ {
			fchan <- uint(rand.Intn(5))
		}
		close(fchan)
	}()

	go func() {
		defer wg.Done()
		for value := range fchan {
			a := float64(value * value * value)
			schan <- a
		}
		close(schan)
	}()

	go func() {
		defer wg.Done()
		for value := range schan {
			fmt.Println(value)
		}
	}()

	wg.Wait()
}
