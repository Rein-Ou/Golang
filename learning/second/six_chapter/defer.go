package main

import (
	"fmt"
	"log"
)

func main() {
	a()
	f()
	fmt.Println(test())
	fmt.Println(f2())

}

func f2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

//defer disconnectFromDB()
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//多defer的时候逆序执行（类似栈）
func f() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\t", i)
	}

}

func test() (i int) {
	i = 0
	defer log.Print("deferred")
	return

}
