package main

import "fmt"

func main() {
	array := new([15]int)
	val_array(array)
	fmt.Println(*array)
}

func val_array(arr *[15]int) {
	lens := len(*arr)
	for i := 0; i < lens; i++ {
		(*arr)[i] = i
	}
}
