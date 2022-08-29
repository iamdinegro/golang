// Solutions for the problem 1007 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int 
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)

	diferenca := (a * b - c * d)
	fmt.Printf("DIFERENCA = %d\n", diferenca)
}