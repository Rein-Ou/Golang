package main

import "fmt"

func main() {
	var a int = 1
	DoSomeThing(&a)
	DoSomeThing2(a)
	fmt.Println(ReturnMulti(5))
	var result int
	Multiply(5, 6, &result)
	fmt.Println(result)
}

func DoSomeThing(a *int) {
	b := a
	fmt.Println(b)
}

func DoSomeThing2(a int) {
	b := &a
	fmt.Println(b)
}

func ReturnMulti(input int) (x1 int, x2 int) {
	x1 = input
	x2 = input * 2
	return
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

//
//func varofinterface(values ...interface{}) {
//	for _, value := range values {
//		switch v := value.(type) {
//		case int:...
//		case float:...
//		}
//	}
//}
