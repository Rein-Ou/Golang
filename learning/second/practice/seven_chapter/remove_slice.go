package main

import "fmt"

func main() {
	stringSlice := []string{"A", "B", "C", "D", "E"}
	stringSlice = RemoveStringSlice(stringSlice, 1, 2)
	fmt.Println(stringSlice)
}

func RemoveStringSlice(Slice []string, start, end int) []string {
	newSlice := make([]string, len(Slice)-(end-start))
	at := copy(newSlice, Slice[:start])
	copy(newSlice[at:], Slice[end:])
	return newSlice
}
