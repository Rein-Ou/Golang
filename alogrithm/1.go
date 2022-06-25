package main

import "fmt"

func duplicateZeros(arr []int) {
	var temp []int
	for _, v := range arr {
		if v != 0 {
			temp = append(temp, v)
		} else {
			temp = append(temp, 0)
			temp = append(temp, v)
		}
	}

	copy(arr, temp)
	arr = append(arr, temp[len(arr):]...)
}

func main() {
	arry := []int{1, 2, 5, 7, 0, 4, 0, 1}
	duplicateZeros(arry)
	fmt.Println(arry)
}
