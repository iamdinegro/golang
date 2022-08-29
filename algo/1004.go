// Solutions for the problem 1004 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var valor1 int 
	var valor2 int 
	fmt.Scanln(&valor1)
	fmt.Scanln(&valor2)
	prod := valor1 * valor2


	fmt.Printf("PROD = %d\n", prod)
}