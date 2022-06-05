package main

import (
	"Golang/src/practice/nine_chapter/fibo/fibo"
	"fmt"
)

var nextFibo int
var op string

func main() {
	op = "+"
	callFunction()
	fmt.Println("change operation to *")
	nextFibo = 0
	op = "*"
	callFunction()

}

func callFunction() {
	nextOperation()
	fmt.Println(">>>>>>>>>>")
	nextOperation()
	fmt.Println(">>>>>>>>>>")
	nextOperation()
	fmt.Println(">>>>>>>>>>")
}

func nextOperation() {
	fmt.Printf("fibonacci(%d),operation (%s),result=%d\n", nextFibo, op, fibo.Fibonacci(op, nextFibo))
	nextFibo++
}
