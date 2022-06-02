package main

import "fmt"

func main() {
	var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd'}
	fmt.Println(string(uniq(arr)))
}

func uniq(arr []byte) (result []byte) {
	arru := make([]byte, len(arr))
	ixu := 0
	tmp := byte(0)
	for _, val := range arr {
		if val != tmp {
			arru[ixu] = val
			ixu++
		}
		tmp = val
	}
	return arru
}
