package main

import "fmt"

func main() {
	//在if开头初始化变量，作用域只在if结构中
	if i := 1; i == 2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	switch result := 8; {
	case result == 0:
	case result == 1:
	case result == 8:
		fmt.Println(8)
	}

}
