package main

import "fmt"

func main() {
	varargs("ass", "ssd", "ddd")
}

func varargs(args ...string) {
	for _, val := range args {
		print(val)
	}
}

func print(args string) {
	fmt.Println(args)
}
