package main

import (
	"fmt"
)

func main() {
	var intChan = make(chan uint)

	go func() {
		defer close(intChan)
		for i := 0; i < 10; i++ {
			intChan <- uint(i)
		}
	}()

	cubed := Cube(intChan)
	convertChan := FloatConverter(cubed)

	for num := range convertChan {
		fmt.Println(num)
	}

}

func Cube(inCh <-chan uint) <-chan uint {
	outCh := make(chan uint)
	go func() {
		defer close(outCh)
		for value := range inCh {
			outCh <- value * value * value
		}
	}()
	return outCh
}

func FloatConverter(inCh <-chan uint) chan float64 {
	outCh := make(chan float64)
	go func() {
		defer close(outCh)
		for value := range inCh {
			outCh <- float64(value)
		}
	}()
	return outCh
}
