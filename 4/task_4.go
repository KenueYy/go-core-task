package main

import (
	"fmt"
)

var (
	slice1 []string = []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 []string = []string{"banana", "date", "fig"}
)

func main() {
	fmt.Print(FilterSlice(slice1, slice2))
}

func FilterSlice(slice1, slice2 []string) []string {

	if len(slice1) == 0 {
		return slice1
	}

	inSecondMap := make(map[string]int)
	resultSlice := make([]string, 0, len(slice1))

	for _, key := range slice2 {
		inSecondMap[key] = 0
	}

	for _, key := range slice1 {
		if _, found := inSecondMap[key]; !found {
			resultSlice = append(resultSlice, key)
		}
	}

	return resultSlice
}
