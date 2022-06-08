package main

import (
	"Golang/src/learning/second/chapter_eleven/stack"
	"fmt"
)

func main() {
	var st stack.Stack
	st.Push("12")
	st.Push(1)
	st.Push(1.33)
	st.Push([]string{"2112", "122312"})
	for {
		item, err := st.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

}
