package main

import "fmt"

var a, b, c int

func main() {
	//var v_name v_type
	var age int

	age = 123
	fmt.Println(age)
	fmt.Println("Hello,World")
	fmt.Println("Hello" + ",World")
	//var v_name=value 根据值自行判断变量类型
	var target = "I'm %d years old"
	fmt.Println(fmt.Sprintf(target, age))
	//v_name:=value 相当于 var v_name v_type ; v_name=value;
	Myvalue := "yes"
	fmt.Println(Myvalue)
	vname1, vname2, vname3 := 1, "rein", 23.4
	fmt.Println(vname1, vname2, vname3)
	fmt.Println(a, b, c)
}

// \t,\n,\\,\",\r
