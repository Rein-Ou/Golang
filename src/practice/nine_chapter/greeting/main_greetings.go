package main

import (
	"Golang/src/practice/nine_chapter/greeting/greetings"
	"fmt"
)

func main() {
	user := "Rein"
	fmt.Println(greetings.GoodDay(user))
	fmt.Println(greetings.GoodNight(user))
	if greetings.IsAM() {
		fmt.Println("早上")
	} else if greetings.IsAfternoon() {
		fmt.Println("下午")
	} else {
		greetings.IsEvening()
		fmt.Println("晚上")
	}
}
