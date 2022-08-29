// Solutions for the problem 1012 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	var pi float64 = 3.14159
	
	fmt.Scanln(&a, &b, &c)

	area_triangulo := a * c / 2
	area_circulo := pi * (c * c)
	area_trapezio := (a + b) * c / 2
	area_quadrado := b * b
	area_retangulo :=  a * b

	fmt.Printf("TRIANGULO: %.3f\n", area_triangulo)
	fmt.Printf("CIRCULO: %.3f\n", area_circulo)
	fmt.Printf("TRAPEZIO: %.3f\n", area_trapezio)
	fmt.Printf("QUADRADO: %.3f\n", area_quadrado)
	fmt.Printf("RETANGULO: %.3f\n", area_retangulo)

}