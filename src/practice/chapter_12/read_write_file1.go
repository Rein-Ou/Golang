package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	inputFile := ".\\src\\practice\\chapter_12\\products.csv"
	outputFile := ".\\src\\practice\\chapter_12\\products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error:%s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
