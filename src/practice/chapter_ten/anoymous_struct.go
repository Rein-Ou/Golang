package main

import "fmt"

type anonymous struct {
	num float64
	int
	string
}

type A struct{ a int }
type B struct{ a, b int }
type C struct {
	A
	B
}

func main() {
	anonymous := anonymous{1.0, 1, "ssss"}
	fmt.Println(anonymous)
	var c C
	c.A.a = 2

}
