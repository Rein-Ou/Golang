package main

import "fmt"

//在Golang中c=*p++是不合法的
func main() {
	//var p *int = nil
	//*p = 0
	string1 := "good bye"
	var ptr *string = &string1
	fmt.Println(*ptr)

	//不能获取到一个文字或常量的地址

}
