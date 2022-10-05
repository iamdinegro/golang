package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Hello, World!\nArguments: %v\n", args[1:len(args)])
}
