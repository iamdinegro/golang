package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 20 { // Não tem While em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		// Laço infinito
		fmt.Println("Para sempre ...")
		time.Sleep(time.Second * 5) // De 5 em 5 Segundos
	}

}
