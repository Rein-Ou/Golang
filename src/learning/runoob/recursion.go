package main

import (
	"fmt"
)

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	fmt.Println("Please input a number")
	fmt.Scanf("%d", &i)
	fmt.Printf("%d的阶乘是%d", i, Factorial(uint64(i)))
	fmt.Print("\n")
	var j int
	for j = 0; j < 10; j++ {
		fmt.Printf("%d\t", fibonacci(j))
	}
}
