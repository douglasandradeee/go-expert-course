package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	/*
		Este é um exemplo da forma simples de se obter o último elemento de qualquer array.
		Passando a função len que retorna o tamanho do mesmo e subtraindo de 1, como o indicador de posição.
	*/
	fmt.Println(meuArray[len(meuArray)-1])

	//Obtendo o primeiro elemento de um array.
	fmt.Println(meuArray[0])

	/*
		Formas de se percorrer um array.
	*/
	for i, v := range meuArray {
		fmt.Printf("\nNa posição %d temos o elemento: %v", i, v)
	}

}
