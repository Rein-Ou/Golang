package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(add(slice...))
}

func add(nums ...int) (res int) {
	result := 0
	for _, value := range nums {
		result += value
	}
	return result
}
