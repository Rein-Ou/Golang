package main

import "fmt"

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

//返回多参数
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := 1, 2
	var ret int
	//调用函数
	ret = max(a, b)
	fmt.Println(ret)
}
