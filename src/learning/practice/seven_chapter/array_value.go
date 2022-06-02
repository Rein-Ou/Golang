package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i << 1
	}

	arr2 := arr1
	arr2[2] = 200

	for pos, val := range arr1 {
		fmt.Printf("Array1 position:%d ,value:%d\n", pos, val)
	}

	for pos, val := range arr2 {
		fmt.Printf("Array2 positiion:%d ,value:%d\n", pos, val)
	}
}
