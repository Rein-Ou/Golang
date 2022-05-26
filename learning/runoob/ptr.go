package main

import "fmt"

func main() {
	var a int = 10
	var d int = 20
	fmt.Printf("a的地址为%x\n", &a)

	var b *int
	b = &a

	fmt.Println(&b)
	fmt.Println(*b)

	var c **int
	c = &b
	fmt.Println(&c)
	fmt.Println(*c)

	fmt.Println(a, d)
	swap(&a, &d)
	fmt.Println(a, d)
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
