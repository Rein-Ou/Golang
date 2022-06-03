package main

import "fmt"

var week = map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 7: "Sunday"}

func main() {
	fmt.Println(week)
	flagHollyday := false
	for key, value := range week {
		if value == "Thursday" || value == "Hollyday" {
			fmt.Println(value, "is the ", key, "th day in the week")
			if value == "Hollyday" {
				flagHollyday = true
			}
		}
	}
	if !flagHollyday {
		fmt.Println("Hollyday is not a day!")
	}
}
