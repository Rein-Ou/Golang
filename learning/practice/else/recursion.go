package main

import "fmt"

func main() {
	recursion_print(1)
	recursion_printVer2(10)
}

func recursion_print(num int) {
	if num > 10 {
		fmt.Println(num)
	} else {
		recursion_print(num + 1)
		fmt.Println(num)
	}
}

func recursion_printVer2(num int) {
	if num > 1 {
		fmt.Println(num)
		recursion_printVer2(num - 1)
	} else {
		fmt.Println(num)
	}
}
