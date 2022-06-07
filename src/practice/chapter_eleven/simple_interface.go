package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (s *Simple) Get() int {
	return s.i
}

func (s *Simple) Set(num int) {
	s.i = num
}

func fI(it Simpler) int {
	it.Set(5)
	return it.Get()
}

func main() {
	var s Simple
	sr := Simpler(&s)
	fmt.Println(fI(sr))
}
