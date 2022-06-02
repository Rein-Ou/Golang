package main

import "fmt"

func main() {
	const LENGTH = 10
	fmt.Println(LENGTH)
	//iota  a special constant 可以认为是一个可以被编译器修改的常量
	const (
		Unkown = iota
		Female        //iota+=1
		Male          //iota+=1
		d      = "le" //iota+=1
		e             //iota+=1
		f             //iota+=1
		g      = iota
		h
	)
	fmt.Println(Unkown, Female, Male, d, e, f, g, h)
}
