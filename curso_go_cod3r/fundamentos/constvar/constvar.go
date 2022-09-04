package main

import (
	"fmt"
	m "math" // Você pode abreviar o nome do pacote para facilitar nas referências do seu código
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma redudiza para criar uma variável
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área de circunferência é ", area)

	// TODAS VARIÁVEIS DECLARADAS DEVEM SER UTILIZADAS!!
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Epa!"
	fmt.Println(g, h, i)

}
