package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("Sending", 10)
	c <- 10
	fmt.Println("Sent", 10)
}
