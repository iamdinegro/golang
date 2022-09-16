package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 13000.00,
			"Guga Pereira":   8200.00,
		},
		"J": {
			"José João": 1200.00,
		},
		"P": {
			"Pedro Junior": 7200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		// Podemos utilizar outro for para percorrer dentro de funcs
		fmt.Println(letra, funcs)
	}
}
