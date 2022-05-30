package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = Filter(s, even)
	fmt.Println(s)
}

func Filter(s []int, fn func(int) bool) (result []int) {
	var newSlice []int
	for _, val := range s {
		if fn(val) {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

func even(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
