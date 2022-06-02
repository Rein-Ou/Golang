package main

import (
	"fmt"
	"time"
)

func main() {
	result := fibonacci(15)
	start := time.Now()
	for pos, val := range result {
		fmt.Printf("fibonacci(%d) is %d\n", pos, val)
	}
	end := time.Since(start)
	fmt.Printf("Time is %s\n", end)
}

func fibonacci(term int) (result []int) {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1
	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}
