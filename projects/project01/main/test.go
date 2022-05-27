package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("error message")
	if errors.Is(e, e) {
		fmt.Printf("%T", e)
	}
	a, b := 10, 0
	c := a / b

	print(c)
}
