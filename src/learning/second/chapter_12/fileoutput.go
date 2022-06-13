package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile(".\\output.dat", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //os.O_APPEND
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)

	}
	outputWriter.Flush()

}
