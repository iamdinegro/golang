// Solutions for the problem 1014 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var km_percorrido int 
	var litros_gastos float64

	fmt.Scanln(&km_percorrido)
	fmt.Scanln(&litros_gastos)

	consumo_medio := float64(km_percorrido) / litros_gastos

	fmt.Printf("%.3f km/l\n", consumo_medio)
}