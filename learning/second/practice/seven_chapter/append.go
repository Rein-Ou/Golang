package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	slice1 := []byte{2, 3, 4}
	slice2 := []byte{6, 7, 8}
	start := time.Now()
	newSlice := AppendBuffer(slice1, slice2)
	fmt.Println(newSlice, len(newSlice))
	end := time.Since(start)
	fmt.Println("time :", end)

}
func Append(slice, data []byte) []byte {
	lengthNewSlice := len(slice) + len(data)
	capNewSlice := cap(slice)
	if lengthNewSlice > capNewSlice {
		capNewSlice = lengthNewSlice
	}

	newSlice := make([]byte, lengthNewSlice, capNewSlice)

	for sliceKey, sliceItem := range slice {
		newSlice[sliceKey] = sliceItem
	}

	for dataKey, dataItem := range data {
		newSlice[dataKey+len(slice)] = dataItem
	}
	return newSlice
}

func AppendBuffer(slice, data []byte) []byte {
	var buffer = bytes.Buffer{}
	buffer.Write(slice)
	buffer.Write(data)
	return buffer.Bytes()
}
