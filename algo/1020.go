package main

import "fmt"

func main() {
	var idade_em_dias int
	fmt.Scanln(&idade_em_dias)

	idade_em_anos := idade_em_dias / 365
	qnt_em_dias := idade_em_dias % 365
	qnt_meses := qnt_em_dias / 30
	qnt_em_dias = qnt_em_dias % 30
	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", idade_em_anos, qnt_meses, qnt_em_dias)

}
