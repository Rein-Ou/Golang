package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	pos := 0
	for i := 0; i <= 10; i++ {
		start := time.Now()
		pos, result = fibonacci(i)
		delta := time.Since(start)
		fmt.Printf("time:%s\n", delta)
		fmt.Printf("fibonacci(%d) ,Position is: %d and result is %d\n", i, pos, result)
	}

}

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func fibonacci(n int) (pos int, result int) {
	if n <= 1 {
		result = 1
	} else {
		_, pre1 := fibonacci(n - 1)
		_, pre2 := fibonacci(n - 2)
		result = pre2 + pre1
	}
	return n + 1, result
}
