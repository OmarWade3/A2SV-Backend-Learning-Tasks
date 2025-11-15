package main

import (
	"fmt"
)

func total(s []int) int {
	result := 0
	for _, val := range s {
		result += val
	}
	return result
}

func main() {
	// ex1
	s1 := []int{3, 6, 9}
	fmt.Println("the sum of: ", s1, " is: ", total(s1))

	// ex2
	s2 := []int{-3, 0, 3, 5}
	fmt.Println("the sum of: ", s2, " is: ", total(s2))
}
