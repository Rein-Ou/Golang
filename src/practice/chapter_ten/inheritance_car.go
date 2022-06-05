package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Merkel")
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

func (c *Car) Start() {
	fmt.Println("Car Start")
}

func (c *Car) Stop() {
	fmt.Println("Car Stop")
}

func main() {
	c := new(Car)
	c.Start()
	c.Stop()

	m := new(Mercedes)
	m.wheelCount = 4
	m.Start()
	m.Stop()
	fmt.Printf("Wheels %d\n", m.numberOfWheels())
	m.sayHiToMerkel()
}

/**
内嵌将一个已存在类型的字段和方法注入到了另一个类型里：匿名字段上的方法 “晋升” 成为了外层类型的方法。当然类型可以有只作用于本身实例而不作用于内嵌 “父” 类型上的方法，

可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。

*/
