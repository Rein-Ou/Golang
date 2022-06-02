package main

import "fmt"

func main() {
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

	for pos, val := range arrKeyValue {
		fmt.Printf("Person at %d is %s\n", pos, val)
	}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
