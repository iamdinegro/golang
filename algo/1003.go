// Solutions for the problem 1003 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var A int
	var B int

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	soma := A + B

	fmt.Printf("SOMA = %d\n", soma)
}