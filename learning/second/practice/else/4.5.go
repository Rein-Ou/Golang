package main

import "fmt"

//string别名
type Rope string

func main() {
	var mystring Rope = "what"
	fmt.Println(mystring)
}
