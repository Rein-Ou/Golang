package main

import (
	"fmt"
	"math"
)

func ConvertInt64ToInt(i int64) (result int) {
	if i <= math.MaxInt32 && i > math.MinInt32 {
		result = int(i)
		return
	}
	panic(fmt.Sprintf("%d is out of range", i))
}

func IntFromInt64(l int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(l)
	return
}

func main() {
	l := int64(15000)
	erL := int64(math.MaxInt64)
	i, err := IntFromInt64(l)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	i2, err2 := IntFromInt64(erL)
	if err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i2)
	}
}
