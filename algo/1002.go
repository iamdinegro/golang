// Solutions for the problem 1002 from breecrowd
package main 

import (
	"fmt"
)

func main() {
	var n float64 = 3.14159
	var raio float64
	fmt.Scanln(&raio)
	var raioaoquadrado = raio * raio
	var area float64 = n * raioaoquadrado
	fmt.Printf("A=%.4f\n", area)
}