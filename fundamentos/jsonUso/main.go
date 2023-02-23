package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"numero"`
	Saldo  float64 `json:"saldo"`
}

func main() {

	minhaConta := Conta{
		Numero: 1,
		Saldo:  1501999.99,
	}

	// Convertendo objeto/struct em JSON
	minhaContaJSON, err := json.Marshal(minhaConta)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(minhaContaJSON))

	// Convertendo objeto/struct em JSON Identado
	minhaContaJSONIdent, err := json.MarshalIndent(minhaConta, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(minhaContaJSONIdent))

	// os.Stdout - significa que ele vai sair pela saída padrão, ou seja, irá sair no terminal.
	// NewEncoder - cria um enconder pra transformar em json já devolvendo o resulta em arquivo ou no terminal.
	// serve para quando não queremos/precisamos gravar o json em uma variável
	err = json.NewEncoder(os.Stdout).Encode(minhaConta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"numero":2, "saldo": 2501999.99}`)
	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", contaX)

}
