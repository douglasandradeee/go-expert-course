package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	response, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// file, err := os.Create("google.html")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// lenFile, err := file.Write(bodyBytes)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Arquivo criado com sucesso! Tamanho:", lenFile)

	fmt.Println(string(bodyBytes))
}
