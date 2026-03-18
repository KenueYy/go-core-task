package main

import (
	"fmt"
	"math/rand"
)

func main() {
	stopChan := make(chan bool)

	fmt.Println(<-RandomGenerator(stopChan, 10))

	close(stopChan)
}

func RandomGenerator(stop <-chan bool, maxValue int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-stop:
				return
			case out <- rand.Intn(maxValue):
			}
		}
	}()

	return out
}
