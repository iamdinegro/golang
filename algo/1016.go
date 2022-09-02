// Solutions for the problem 1016 from breecrowd
package main


import (
	"fmt"
)

func main() {
	var distancia_km int

	fmt.Scanln(&distancia_km)

	tempo := distancia_km * 2

	fmt.Printf("%d minutos\n", tempo)

}