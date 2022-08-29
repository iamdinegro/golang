// Solutions for the problem 1011 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var pi float64 = 3.14159
	var raio_esfera float64

	fmt.Scanln(&raio_esfera)
	raio_esfera_ao_cubo := raio_esfera * raio_esfera * raio_esfera

	ve := (4.0 / 3) * pi * raio_esfera_ao_cubo
	fmt.Printf("VOLUME = %.3f\n", ve)
}