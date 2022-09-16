package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // Compilador que vai contar

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)

	}
	for _, num := range numeros { // Para ignorar o indice você pode colocar um "_"
		fmt.Println(num)
	}
}
