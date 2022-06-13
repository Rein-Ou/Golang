package main

import (
	"Golang/src/learning/second/chapter_eleven/stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	calc1 := new(stack.Stack)
	fmt.Println("Give a number, an operator(+,-,*,/),or q to stop:")
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Error")
			os.Exit(0)
		}
		token = token[0 : len(token)-2] //remove \r\n
		switch {
		case token == "q":
			fmt.Println("Calculator stopped")
			os.Exit(0)
		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc1.Push(i)
		case token == "+":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)+q.(int))
		case token == "-":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)-q.(int))
		case token == "*":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)*q.(int))
		case token == "/":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)/q.(int))
		default:
			fmt.Println("Invalid input")

		}
	}
}
