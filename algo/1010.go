// Solutions for the problem 1010 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var codigo_peca_1 int
	var numero_peca_1 int 
	var valor_peca_1 float64
	var codigo_peca_2 int
	var numero_peca_2 int 
	var valor_peca_2 float64

	fmt.Scanln(&codigo_peca_1, &numero_peca_1, &valor_peca_1)
	fmt.Scanln(&codigo_peca_2, &numero_peca_2, &valor_peca_2)

	total_peca_1 := float64(numero_peca_1) * valor_peca_1
	total_peca_2 := float64(numero_peca_2) * valor_peca_2
	total_a_pagar := total_peca_1 + total_peca_2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total_a_pagar)
}