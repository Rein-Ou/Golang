package main

import "fmt"

func main() {
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i
	}
	var slice1 []int
	var slice2 []int = make([]int, 5)
	slice3 := make([]int, 3)
	slice1 = []int{1, 2, 3}
	slice2 = array[3:8]
	slice1 = append(slice1, 5)
	fmt.Println(slice1)
	fmt.Println(slice2)
	//长度
	fmt.Println(len(slice3))
	//容量
	fmt.Println(cap(slice3))
	/**
	s[1:]
	s[:4]
	**/

	var numbers []int
	if numbers == nil {
		fmt.Printf("切片是空的")
	}

}
