package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// interação com usuário.
	var num1, num2 int
	fmt.Println("Por favor informe os números que deseja somar:")
	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		log.Fatal("Erro ao capturar números da soma")
	}

	valor, err := sumVariadic(num1, num2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("O resultado da soma é:", valor)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}

// sumVariadic- é praticamente a mesma função de somar, mas escrita de forma variádica. Feita para receber parâmetros indeterminados.
func sumVariadic(numeros ...int) (int, error) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	if total >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return total, nil
}
