package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado..
	fmt.Println("Teste " + string(97)) // 97 aponta para a letra A na tabela ASC

	// int para string
	fmt.Println("Test " + strconv.Itoa(123))

	// String para int

	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro") // Se B for verdadeiro ele printa "Verdadeiro"
	}
}
