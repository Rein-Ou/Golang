package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	defer fmt.Printf("The process id is %v", pid)
}
