package main

import "fmt"

type Triangle struct {
}

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}
type RSimple struct {
	i int
	j int
}

func (rs *RSimple) Get() int {
	return rs.j
}

func (rs *RSimple) Set(num int) {
	rs.j = num
}

func (s *Simple) Get() int {
	return s.i
}

func (s *Simple) Set(num int) {
	s.i = num
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(10)
		return it.Get()
	default:
		return -1
	}
}

func main() {
	var s Simple
	sr := Simpler(&s)
	fmt.Println(fI(sr))
	var rss RSimple
	sr = Simpler(&rss)
	fmt.Println(fI(sr))
}
