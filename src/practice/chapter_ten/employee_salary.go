package main

import (
	"fmt"
	"time"
)

type employee struct {
	salary float64
}

func (e *employee) giveRaise(percent float64) {
	e.salary += e.salary * (percent / 100)
}

func main() {
	emp := &employee{1000}
	emp.giveRaise(51)
	fmt.Println(emp.salary)
	fmt.Println(time.Now().String()[0:3])
}
