package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 10, 20)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	for pos, val := range slice1 {
		fmt.Printf("Slice at %d is %d\n", pos, val)
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
