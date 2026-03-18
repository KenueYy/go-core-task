package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	stopChan := make(chan bool)
	randChan1 := RandomGenerator(stopChan, 2)
	randChan2 := RandomGenerator(stopChan, 10)
	randChan3 := RandomGenerator(stopChan, 100)

	out := MergeChannel(randChan1, randChan2, randChan3)

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

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

func MergeChannel(chs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))
	for _, ch := range chs {
		go func(c <-chan int) {
			defer wg.Done()

			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
