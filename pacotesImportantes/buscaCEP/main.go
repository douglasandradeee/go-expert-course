package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

const URL_VIA_CEP string = "http://viacep.com.br/ws/"

func main() {
	// captura os dados digitados no terminal e os imprime.
	for _, cep := range os.Args[1:] {
		response, err := http.Get(URL_VIA_CEP + cep + "/json/")
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

		file, err := os.Create(strings.ToLower(MyCEP.Cep + "_" + strings.ReplaceAll(MyCEP.Localidade, " ", "_") + ".txt"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao criar o arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s, Logradouro: %s, Bairro: %s, DDD: %s", MyCEP.Cep, MyCEP.Localidade, MyCEP.Uf, MyCEP.Logradouro, MyCEP.Bairro, MyCEP.Ddd))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao escrever no arquivo: %v\n", err)
		}
		fmt.Printf("Arquivo %s criado com sucesso.\n", file.Name())
		fmt.Printf("%+v", MyCEP)
	}

}
