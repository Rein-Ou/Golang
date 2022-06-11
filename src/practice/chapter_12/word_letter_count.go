package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words, runes, lines int

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input something:")
	for {
		inputStrings, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error")
			break
		} else if inputStrings == "S\r\n" {
			fmt.Println("lines:", lines)
			fmt.Println("words", words)
			fmt.Println("runes", runes)
			os.Exit(0)
		}
		Counts(inputStrings)

	}

}

func Counts(inputStrings string) {
	runes += len(inputStrings) - 2
	words += len(strings.Fields(inputStrings))
	lines += 1
}
