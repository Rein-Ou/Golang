package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") //-n 初始值为false，帮助信息为usage
//var l = flag.Int("w", 0, "we")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	//flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine {
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
