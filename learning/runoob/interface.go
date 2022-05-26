package main

import "fmt"

//定义接口
type Phone interface {
	call()
}

type NokiaPhone struct {
}

type iPhone struct {
}

//实现接口
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you!")
}

func (iPhone iPhone) call() {
	fmt.Println("I am iPhone,I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	phone = new(iPhone)
	phone.call()
}
