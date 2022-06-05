package main

import (
	"Golang/src/practice/chapter_ten/rectangle/rectangle"
	"fmt"
)

func main() {
	r1 := rectangle.NewRectangle(5, 6)
	r2 := rectangle.NewRectangle(5, 6)
	r2.SetLength(6)
	fmt.Println(r1.GetWidth())
	fmt.Println(rectangle.Area(*r1))
	fmt.Println(r2)
	fmt.Println(rectangle.Primeter(*r2))
}
