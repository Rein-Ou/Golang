package main

import (
	"fmt"
	"unicode"
)

func main() {
	var ch int = '\u0041' //'A'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", rune(ch), ch2, ch3) //integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)       //character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3)       //UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3)       //UTF-8 code point
	//type rune=int32
	fmt.Printf("%t\n", unicode.IsLetter(rune(ch)))
	fmt.Printf("%t\n", unicode.IsDigit(rune(ch)))
	fmt.Printf("%t\n", unicode.IsSpace(rune(ch)))

}
