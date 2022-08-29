// Solutions for the problem 1013 from breecrowd
package main 

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int

	fmt.Scanln(&a, &b, &c)
	abs_ab := math.Abs(float64(a) - float64(b))
	maiorAB := (a + b + int(abs_ab)) / 2
	if (c > maiorAB) {
		fmt.Printf("%d eh o maior\n", c)
	} else {
		fmt.Printf("%d eh o maior\n", maiorAB)
	}
	
}