// filecopy.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := CopyFile("target.txt", "source.txt")
	if err == nil {
		fmt.Println("Copy done!")
	} else {
		fmt.Println(err)
	}
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
