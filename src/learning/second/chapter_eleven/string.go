package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}
type aaa struct {
	ss int
}

func (s *aaa) String() string {
	return strconv.Itoa(s.ss)
}

func main() {
	var v aaa
	v.ss = 123
	ss := Stringer(&v)
	if sv, ok := ss.(Stringer); ok {
		fmt.Printf("v implemetns String():%s\n", sv.String())
	}

}
