package main

import "fmt"

var i int = 0

//break,continue,goto
func main() {
	for li() {
		fmt.Println("死循环了")
	}
}

func li() bool {
	if i == 2 {
		return false
	}
	i += 1
	return true
}
