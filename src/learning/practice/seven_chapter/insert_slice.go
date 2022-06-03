package main

import "fmt"

func main() {
	slFrom := []string{"12", "34"}
	insertion := []string{"56", "78"}
	result := insert(slFrom, insertion, 1)
	fmt.Println(result)
}

func insert(slFrom, insert []string, index int) []string {
	result := make([]string, len(slFrom)+len(insert))
	at := copy(result, slFrom[:index])
	at += copy(result[at:], insert)
	copy(result[at:], slFrom[index:])
	return result
}
