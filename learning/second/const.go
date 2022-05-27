package main

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	fmt.Println(GB)
	const hardEight = (1 << 100) >> 97
	fmt.Println(hardEight)

	const (
		Sunday = iota
		Monday
		Tuesday
	)
	//每一次遇到const关键字 iota就重置为0
	const (
		Wednesday = iota
		Thursday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday)
}
