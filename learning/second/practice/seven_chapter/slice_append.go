package main

import "fmt"

func main() {
	slice1 := make([]int, 2)
	slice1 = func(s []int, factor int) []int {
		newSlice := make([]int, len(s)*factor)
		copy(newSlice, s)
		s = newSlice
		return s
	}(slice1, 2)
	fmt.Println(slice1, len(slice1))
}
