package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting the program")
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("Runtime panic :%v\n", err)
		}
	}()
	log.Panic("A severe error occurred:stopping the program!")
}
