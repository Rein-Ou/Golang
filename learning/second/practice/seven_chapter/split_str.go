package main

import (
	"fmt"
	"time"
)

func main() {
	string := "Google"
	slice := []byte(string)
	fmt.Println(slice)
	fmt.Printf("%s\n", slice)
	fmt.Println(split(string, 3))
	fmt.Println(string[len(string)/2:] + string[:len(string)/2])
	fmt.Println(reverse(string))
}

func split(str string, i int) (pre, suffix string) {
	stringSlice := []byte(str)
	pre = string(stringSlice[:i])
	suffix = string(stringSlice[i:])
	return
}

func reverse(str string) (result string) {
	start := time.Now()
	temp := make([]byte, len(str))
	pre := str[:len(str)/2]
	suffix := str[len(str)/2:]
	for i := 0; i < len(str)/2; i++ {
		temp = append(temp, suffix[len(str)/2-i-1])
	}
	for i := 0; i < len(str)/2; i++ {
		temp = append(temp, pre[len(str)/2-i-1])
	}

	end := time.Since(start)
	fmt.Println(end)
	return string(temp)
}
