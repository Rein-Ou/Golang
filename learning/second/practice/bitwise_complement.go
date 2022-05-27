package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ {
		fmt.Printf("%d的补码为%b\n", i, ^i)
	}
}
