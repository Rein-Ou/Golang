package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad Call\r\n")
}

func main() {
	fmt.Println("Calling")
	test()
	fmt.Println("Test Completed")
}
