package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unsafe"
)

var title []string
var prices []float64
var quantities []int

func main() {
	path, _ := filepath.Abs("./src/practice/chapter_12/products.csv")
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	result, _ := reader.ReadAll()

	for _, value := range result { //[][]string
		for i := 0; i < len(value); i++ { //[]string
			switch i {
			case 0:
				title = append(title, value[i])
			case 1:
				price, _ := strconv.ParseFloat(value[i], int(unsafe.Sizeof(value[i])))
				prices = append(prices, price)
			case 2:
				quantity, _ := strconv.Atoi(value[i])
				quantities = append(quantities, quantity)
			}
		}
	}
	fmt.Println(title)
	fmt.Println(prices)
	fmt.Println(quantities)
}
