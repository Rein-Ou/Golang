package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := ver1(5)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func ver1(num float64) (result float64, err error) {
	if num < 0 {
		err = errors.New("不能为负数")
	} else {
		result = math.Sqrt(num)
	}
	return
}
func ver2(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("不能为负数")
	} else {
		return math.Sqrt(num), nil
	}
}
