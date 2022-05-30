package main

import "fmt"

func main() {
	array := []int{1, 8, 3, 68, 93, 2, 7, 1}
	fmt.Println(array)
	bubble(array)
	fmt.Println(array)
}

//冒泡排序
func bubble(arr []int) {
	for pass := 1; pass < len(arr); pass++ {
		for i := 0; i < len(arr)-pass; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}
