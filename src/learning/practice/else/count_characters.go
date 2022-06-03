package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	string1 := "asSASA ddd dsjkdsjs dk"
	string2 := "asSASA ddd dsjkdsjsこん dk"

	fmt.Printf("%d\n", len(string1))
	fmt.Printf("%d\n", count_string(string1))
	//bytes
	fmt.Printf("%d\n", len(string2))
	//runes
	fmt.Printf("%d\n", count_string(string2))

}

func count_string(mystring string) int {
	return utf8.RuneCountInString(mystring)
}
