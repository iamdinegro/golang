// Solutions for the problem 1009 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var nome string 
	var salario float64
	var vendas float64

	fmt.Scanln(&nome)
	fmt.Scanln(&salario)
	fmt.Scanln(&vendas)

	pagamento_adicional := vendas * 15 / 100 
	pagamento_total :=  pagamento_adicional + salario
	fmt.Printf("TOTAL = R$ %.2f\n", pagamento_total)
}