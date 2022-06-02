package main

import (
	"container/list"
	"fmt"
)

func main() {
	List := list.New()
	List.PushFront(101)
	List.PushFront(102)
	List.PushFront(103)
	for i := List.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
