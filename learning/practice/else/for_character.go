package main

import "fmt"

func main() {
	for_character()
	for_character2()
}

func for_character() {
	for i := 1; i <= 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
}

func for_character2() {
	var str string = "G"
	for i := 1; i <= 25; i++ {
		fmt.Println(str)
		str += "G"
	}
}
