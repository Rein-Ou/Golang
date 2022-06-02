package main

import "fmt"

func main() {
	//e := errors.New("error message")
	//if errors.Is(e, e) {
	//	fmt.Printf("%T", e)
	//}
	//a, b := 10, 0
	//c := a / b
	//
	//print(c)

	//var x = []int{2, 3, 4, 5, 6}[2:3]
	//fmt.Println(len(x), cap(x))

	// make([]type,len,cap)==new([100]int)[0:50]
	//var slice []int = make([]int, 5)
	items := [...]int{10, 20, 30, 40, 50}

	for _, item := range items {
		item *= 2
	}
	fmt.Println(items)
	for i := 0; i < len(items); i++ {
		items[i] *= 2
	}
	fmt.Println(items)
}
