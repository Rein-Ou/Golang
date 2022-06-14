package main

import (
	"Golang/src/practice/nine_chapter/even/even"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the interger %d even? %v\n", i, even.Even(i))
	}
}
