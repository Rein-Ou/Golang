package main

import "fmt"

var fib [50]uint64

func main() {
	fib[0] = 1
	fib[1] = 1
	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println(fib)
}
