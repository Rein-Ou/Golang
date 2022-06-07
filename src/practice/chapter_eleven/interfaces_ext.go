package main

import "fmt"

type Square struct {
	side float32
}

type Triangle struct {
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PerInterface interface {
	Perimeter() float32
}

func main() {
	tri := &Triangle{3, 4}
	fmt.Println(AreaInterface(tri).Area())
	sq := &Square{4}
	fmt.Println(PerInterface(sq).Perimeter())
}

func (t *Triangle) Area() float32 {
	return t.base * t.height * 0.5
}

func (s *Square) Perimeter() float32 {
	return s.side * 4
}
