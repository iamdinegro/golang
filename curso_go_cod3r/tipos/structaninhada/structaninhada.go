package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10}, // Use a estrutura assim para não se perder no código
			item{2, 100, 3.56},
			item{5, 5, 7.9},
		},
	}

	fmt.Printf("Valor total do pedido é: %.2f", pedido.valorTotal())
}
