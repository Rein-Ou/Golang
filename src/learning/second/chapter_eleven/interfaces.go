package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (sq *Square) Area() float32 { //实现了Shaper接口
	return sq.side * sq.side
}

func (rec *Rectangle) Area() float32 {
	return rec.length * rec.width
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	rec := new(Rectangle)
	rec.length = 4
	rec.width = 5
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Printf("The square has a area:%f\n", areaIntf.Area())
	areaIntf = rec
	fmt.Printf("The rectangle has a area:%f\n", areaIntf.Area())

	shapes := []Shaper{sq1, rec}

	for _, value := range shapes {
		fmt.Println(value.Area())
	}

}
