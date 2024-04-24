package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	ApiUrl string
	Data   interface{}
	Error  error
}

func makeRequest(url string, ch chan<- Response) {
	start := time.Now()
	client := http.Client{Timeout: time.Second * 2}
	response, err := client.Get(url)
	if err != nil {
		ch <- Response{ApiUrl: url, Error: err}
		return
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		ch <- Response{ApiUrl: url, Error: err}
		return
	}

	var data interface{}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		ch <- Response{ApiUrl: url, Error: err}
		return
	}

	elapsed := time.Since(start)
	ch <- Response{ApiUrl: url, Data: data}
	fmt.Println("Tempo de requisição:", elapsed, "url:", url)
}

func main() {
	var cep string
	fmt.Println("Digite o CEP:")
	fmt.Scanln(&cep)

	if len(cep) == 8 {
		cep = cep[:5] + "-" + cep[5:]
	}

	ch := make(chan Response, 2)

	go makeRequest("https://viacep.com.br/ws/"+cep+"/json/", ch)
	go makeRequest("https://cdn.apicep.com/file/apicep/"+cep+".json", ch)

	var fasterResponse Response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < 2; i++ {
		select {
		case resp := <-ch:
			if resp.Error == nil || resp.ApiUrl == "" {
				fasterResponse = resp
			}
		case <-ctx.Done():
			resp := <-ch
			fmt.Println("Tempo limite de requisição excedido - url:", resp.ApiUrl)
			break
		}
	}

	if fasterResponse.Error != nil {
		fmt.Println("Error:", fasterResponse.Error)
	} else {
		fmt.Println("Api mais rápida:", fasterResponse.ApiUrl)
		fmt.Println("dados:", fasterResponse.Data)

	}
}
