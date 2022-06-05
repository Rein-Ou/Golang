package main

import (
	"Golang/src/practice/nine_chapter/even/even"
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("数字%d是偶数吗？%t\n", i, even.Even(i))
	}
}
