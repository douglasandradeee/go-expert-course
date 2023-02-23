package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// captura os dados digitados no terminal e os imprimi
	for _, url := range os.Args[1:] {
		println(url)
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao fazer requisição: %v\n", err)
		}
		defer response.Body.Close()

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao ler a resposta: %v\n", err)
		}

		var MyCEP ViaCEP
		err = json.Unmarshal(bodyBytes, &MyCEP)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao fazer o parse da resposta: %v\n", err)
		}
		fmt.Println(MyCEP)
	}

}
