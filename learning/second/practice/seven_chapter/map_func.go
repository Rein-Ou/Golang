package main

import "fmt"

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6}
	mf := func(num int) int {
		return num * 10
	}
	result := mapFunc(mf, list)
	fmt.Println(result)
}

func mapFunc(mf func(int) int, list []int) []int {
	result := make([]int, len(list))
	for ix, val := range list {
		result[ix] = mf(val)
	}
	return result
}
