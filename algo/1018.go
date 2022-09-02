// Solutions for the problem 1018 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var qtdNotas int

	fmt.Scanln(&qtdNotas)
	fmt.Println(qtdNotas)

	fmt.Printf("%d nota(s) de R$ 100,00\n", (qtdNotas / 100))
	aux := qtdNotas % 100

	fmt.Printf("%d nota(s) de R$ 50,00\n", (aux / 50))
	aux = aux % 50

	fmt.Printf("%d nota(s) de R$ 20,00\n", (aux / 20))
	aux = aux % 20

	fmt.Printf("%d nota(s) de R$ 10,00\n", (aux / 10))
	aux = aux % 10

	fmt.Printf("%d nota(s) de R$ 5,00\n", (aux / 5))
	aux = aux % 5

	fmt.Printf("%d nota(s) de R$ 2,00\n", (aux / 2))
	aux = aux % 2

	fmt.Printf("%d nota(s) de R$ 1,00\n", (aux / 1))
	aux = aux % 1
}
