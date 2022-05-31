package main

import "fmt"

func main() {
	fmt.Println(ver1(10, 5))
	fmt.Println(ver2(10, 5))
}

func ver1(num1 int, num2 int) (int, int, int) {
	return num1 + num2, num1 * num2, num1 - num2
}

func ver2(num1 int, num2 int) (add int, mul int, sub int) {
	add = num1 + num2
	mul = num1 * num2
	sub = num1 - num2
	return
}
