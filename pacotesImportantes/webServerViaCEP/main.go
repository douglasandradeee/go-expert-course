package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	API_PORT    string = ":8080"
	URL_VIA_CEP string = "https://viacep.com.br/ws/"
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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(API_PORT, nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	dataCep, err := BuscaCEP(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataCep)
}

func BuscaCEP(cep string) (*ViaCEP, error) {
	response, err := http.Get(URL_VIA_CEP + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var dataCep ViaCEP
	err = json.Unmarshal(bodyBytes, &dataCep)
	if err != nil {
		return nil, err
	}

	return &dataCep, nil
}
