package main

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

type Endereco struct {
}

func (c *Cliente) Desativar() {
	c.Ativo = false
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e *Empresa) Desativar() {}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	douglas := Cliente{
		Nome:  "Douglas",
		Idade: 30,
		Ativo: true,
	}
	minhaEmmpresa := Empresa{Nome: "DEV.GO"}
	Desativacao(&minhaEmmpresa)
	Desativacao(&douglas)
}
