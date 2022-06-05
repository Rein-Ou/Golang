package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	fmt.Println(ms)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	fmt.Printf("The int is:%d\n", ms.i1)
	fmt.Printf("The float is:%f\n", ms.f1)
	fmt.Printf("The string is:%s\n", ms.str)
	fmt.Println(ms)
	ptr := struct1{1, 2.5, "ee"}
	pptr := &struct1{1, 2.5, "ee"} //==new(struct1)
	fmt.Println(ptr, ptr.i1)
	fmt.Println(pptr.i1)
}
