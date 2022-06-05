package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func UpPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	//value
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	UpPerson(&pers1)
	fmt.Println(pers1)
	//pointer,没有像 C++ 中那样需要使用 -> 操作符，Go 会自动做这样的转换。
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" //valid ,==pers1.lastName
	UpPerson(pers2)
	fmt.Println(pers2)
	//literal
	pers3 := &Person{"Chris", "Woodward"}
	UpPerson(pers3)
	fmt.Println(pers3)
}
