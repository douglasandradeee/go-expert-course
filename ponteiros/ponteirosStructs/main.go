package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo float64
}

func NewConta() *Conta {
	return &Conta{saldo: 0.00}
}

func (co *Conta) simular(valor float64) float64 {
	co.saldo += valor
	fmt.Println(co.saldo)
	return co.saldo
}

func (c *Cliente) andou() {
	c.nome = "Douglas Andrade"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	douglas := Cliente{
		nome: "Douglas",
	}
	douglas.andou()
	fmt.Printf("O valor da strcut com nome %v\n", douglas.nome)

	conta := Conta{
		saldo: 100.00,
	}
	conta.simular(200.00)
	fmt.Println(conta.saldo)

}
