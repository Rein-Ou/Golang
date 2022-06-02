package main

import "fmt"

func main() {
	var a, b = 21, 10
	var c int
	c = a + b
	fmt.Printf("1-c:%d\n", c)
	c = a - b
	fmt.Printf("2-c:%d\n", c)
	c = a * b
	fmt.Printf("3-c:%d\n", c)
	//取整
	c = a / b
	fmt.Printf("4-c:%d\n", c)
	//取余
	c = a % b
	fmt.Printf("5-c:%d\n", c)
	a++
	fmt.Printf("6-c:%d\n", a)
	a = 21
	a--
	fmt.Printf("7-c:%d\n", a)

	// <,>,<=,>=
	if a == b {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	d, f := true, false
	// &&：and, ||：or,!:not
	if d && f {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	/**
		位运算
		&:与,|:或,^:异或,>>:右移,<<:左移
	**/

	var ptr *int = &a
	fmt.Println(*ptr)
	fmt.Println(&a)
}
