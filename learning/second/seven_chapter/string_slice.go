package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "\u00ff\u754c"
	c := []int32(s)
	for i, v := range c {
		fmt.Printf("%d:%c \n", i, v)
	}
	fmt.Println(utf8.RuneCountInString(s))
}
