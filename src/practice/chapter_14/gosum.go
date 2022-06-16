package main

import "fmt"

func main() {
	result := make(chan int)
	go sum(1, 2, result)
	fmt.Println("1 + 2 = ", <-result)
}

func sum(a, b int, result chan int) {
	result <- a + b
}
