package main

import "fmt"

func main() {
	for_loop()
	for_loop2()
}

func for_loop() int {
	var j int = 0
	for i := 0; i < 15; i++ {
		j++
	}
	fmt.Println(j)
	return j
}

//goto方式
func for_loop2() int {
	i, j := 15, 0
loop:
	if i == 0 {
		fmt.Println(j)
		return j
	}
	j++
	i--
	goto loop
}
