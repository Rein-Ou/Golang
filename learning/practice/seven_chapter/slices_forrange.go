package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 4)

	for i := 1; i <= 4; i++ {
		slice1[i-1] = i
	}

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is:%d\n", ix, value)
	}

	for ix, _ := range slice1 {
		slice1[ix] = ix + 5
	}
	fmt.Println(slice1)
}
