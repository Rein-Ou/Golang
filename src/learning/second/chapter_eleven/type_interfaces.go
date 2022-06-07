package main

import "fmt"

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is:%T\n", t)
	}

}
