package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	file, err := os.OpenFile("vcard.gob", os.O_RDONLY, 0666)
	if err != nil {
		log.Println("Error")
		return
	}
	defer file.Close()
	dec := gob.NewDecoder(file)
	var value VCard
	err = dec.Decode(&value)
	if err != nil {
		log.Println("Error")
		return
	}
	fmt.Println(value)

}
