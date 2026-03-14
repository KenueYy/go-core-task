package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := originalSlice(10)
	c := copySlice(s)
	fmt.Println(s)
	fmt.Println(sliceExample(s))
	fmt.Println(addElement(s, 10))
	fmt.Println(c)
	fmt.Println(removeElement(s, 3))

}

func originalSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(100)
	}
	return s
}

func sliceExample(slice []int) []int {
	result := make([]int, 0, len(slice))
	for _, value := range slice {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func addElement(slice []int, number int) []int {
	return append(slice, number)
}

func copySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)

	return result
}

func removeElement(slice []int, pos int) []int {
	return append(slice[:pos], slice[pos+1:]...)
}
