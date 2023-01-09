package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	cliente1 := Cliente{Nome: "Gabriela", Idade: 29, Ativo: true}

	fmt.Printf("Nome: %s Idade: %d Ativo: %t", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
	fmt.Printf("\ncliente_1: %+v", cliente1)

	cliente1.Ativo = false
	fmt.Printf("\ncliente_1: %+v", cliente1)

}
