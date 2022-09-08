package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 5:
		return "C"
	default:
		return "D"
	}
}

func main() {
	fmt.Println(notaParaConceito(3.5))
}
