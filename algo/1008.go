// Solutions for the problem 1008 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var funcionario_num, horas_trabalhadas int
	var salario_por_hora float64
	
	fmt.Scanln(&funcionario_num)
	fmt.Scanln(&horas_trabalhadas)
	fmt.Scanln(&salario_por_hora)

	salario := float64(horas_trabalhadas) * salario_por_hora
	fmt.Printf("NUMBER = %d\n", funcionario_num)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}