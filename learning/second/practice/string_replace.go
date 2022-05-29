package main

import (
	"fmt"
	"strings"
)

func main() {
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return '?'
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "什么东西哎"))
}
