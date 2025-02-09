package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	apiKey := "YOUR_SECRET_OPENAI_API_KEY"

	//mensagem para enviar ao chat gpt-3
	msg := "Me envie os números sorteados da mega sena da virada dos últimos 20 anos"

	reqBody, err := json.Marshal(map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"prompt":      msg,
		"max_tokens":  4000,
		"temperature": 1.0,
	})

	if err != nil {
		fmt.Println("Erro ao criar o corpo da requisição: ", err.Error())
		return
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Erro ao criar a requisição: ", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a requisição: ", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta: ", err.Error())
		return
	}

	fmt.Println(string(body))
}
