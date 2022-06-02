package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{
		"beer":   "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé"}

	sdrinks := make([]string, len(drinks))
	ix := 0
	fmt.Println("Drinks are available")
	for key := range drinks {
		sdrinks[ix] = key
		ix++
	}
	fmt.Println(sdrinks)

	for key, value := range drinks {
		fmt.Println(key, value)
	}

	sort.Strings(sdrinks)
	fmt.Println(sdrinks)

	for _, value := range sdrinks {
		fmt.Println(value, drinks[value])
	}

}
