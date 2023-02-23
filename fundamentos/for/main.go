package main

import "fmt"

func main() {
	numeros := []string{"um", "dois", "três"}

	for k, v := range numeros {
		fmt.Println(k, v)
	}

	i := 0

	for i < 10 {
		println(i)
		i++
	}
	//loop infinito, geralmente utilizado para consumir mensagens de Filas etc.
	for {

	}

}
