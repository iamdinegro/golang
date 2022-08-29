// Solutions for the problem 1006 from breecrowd
package main

import (
	"fmt"
)

func main() {
	var a, b, c, media float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	media = ((2 * a) + (3 * b) + (c * 5)) / 10

	fmt.Printf("MEDIA = %.1f\n", media)
}