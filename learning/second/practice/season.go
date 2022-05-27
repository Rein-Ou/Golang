package main

import "fmt"

func main() {
	fmt.Println(Season(0))
}

func Season(month int) string {
	if month <= 0 || month > 12 {
		return "month error"
	}
	var seasonString string
	switch month {
	case 12, 1, 2:
		seasonString = "Winter"
	case 3, 4, 5:
		seasonString = "Spring"
	case 6, 7, 8:
		seasonString = "Summer"
	case 9, 10, 11:
		seasonString = "Autumn"
	}
	return seasonString

}
