package main

import "fmt"

func main() {
	var mapLit map[string]int

	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	//不要使用 new，永远用 make 来构造 map
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	for key, value := range mapCreated {
		fmt.Printf("%s:%f\n", key, value)
	}
}
