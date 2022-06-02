package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)

	fmt.Println(sl_to)
	fmt.Println(n)
}
