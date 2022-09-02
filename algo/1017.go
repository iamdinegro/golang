// Solutions for the problem 1017 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var consumo_medio int = 12
	var tempo_viagem, velocidade_media int 
	
	fmt.Scanln(&tempo_viagem)
	fmt.Scanln(&velocidade_media)

	litros_necessarios := float64(tempo_viagem) * float64(velocidade_media) / float64(consumo_medio)
	fmt.Printf("%.3f\n", litros_necessarios)
}