package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Print(IntersectionCheck(a, b))
}

func IntersectionCheck(slice1, slice2 []int) (bool, []int) {
	rm := make(map[int]int, len(slice1))
	var resultSlice []int

	for _, key := range slice1 {
		if _, found := rm[key]; !found {
			rm[key] = 0
		}
	}

	for _, key := range slice2 {
		if _, found := rm[key]; !found {
			rm[key] = 0
		} else {
			resultSlice = append(resultSlice, key)
		}
	}

	if len(resultSlice) > 0 {
		return true, resultSlice
	}

	return false, make([]int, 0)
}
