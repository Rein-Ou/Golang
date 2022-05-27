package main

import "fmt"

func main() {
	//``中的为非解释字符串
	mystring := `This is a raw string \n`
	newstring := `This is a raw strisd \n`
	fmt.Println(mystring)
	//可以使用比较运算符在内存中按字节比较实现字符串的对比,比较的是字符串内容
	fmt.Printf("%t\n", mystring <= newstring)
	fmt.Println(len(mystring))

	for pos, char := range mystring {
		fmt.Printf("%d-%c\n", pos, char)
	}
}
