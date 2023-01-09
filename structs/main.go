package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Endereco
	Ativo bool
}

// Desativar - Método do objeto cliente que permite desativar.
func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("\nCliente %s foi desativado.", c.Nome)
}

func main() {

	cliente1 := Cliente{Nome: "Gabriela", Idade: 29, Ativo: true}

	//fmt.Printf("Nome: %s Idade: %d Ativo: %t", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
	fmt.Printf("\ncliente_1: %+v", cliente1)

	cliente1.Desativar()
	fmt.Printf("\ncliente_1_desa: %+v", cliente1)

	cliente1.Logradouro = "Rua aleatória"
	cliente1.Numero = 15
	cliente1.Cidade = "Salvador"
	cliente1.Estado = "Bahia"

	fmt.Printf("\ncliente_1: %+v", cliente1)

}
