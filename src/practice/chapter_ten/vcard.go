package main

import (
	"fmt"
	"time"
)

type Address struct {
	position string
}

type VCard struct {
	name         string
	pic          string
	address_code int
	birthday     time.Time
	Address      *Address
}

func main() {
	address := &Address{"sssss"}
	asan := &VCard{"AsaMi", "DDDD", 1234, time.Now(), address}
	fmt.Println(asan)
}
