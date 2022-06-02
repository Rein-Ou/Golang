package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Min is %d\n", minSlice(slice))
	fmt.Printf("Max is %d\n", maxSlice(slice))

}

func minSlice(slice []int) (res int) {
	min := slice[0]
	for _, item := range slice {
		if min > item {
			min = item
		}
	}
	return min
}

func maxSlice(slice []int) (res int) {
	max := slice[0]
	for _, item := range slice {
		if max < item {
			max = item
		}
	}
	return max
}
