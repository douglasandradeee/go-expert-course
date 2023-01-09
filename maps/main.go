package main

import "fmt"

func main() {
	salarios := map[string]int{"Douglas": 1000.00, "João": 2000.00, "Maria": 30000.00}
	fmt.Println(salarios["Maria"])

	//delete é utilizado para remover um dado de um map através da chave indicada.
	delete(salarios, "João")
	fmt.Println(salarios)

	// adicionando dado a um map já existente
	salarios["Gabi"] = 900000.00
	fmt.Println(salarios)

	// para percorrer um map nós também utilizamos um "for".
	for nome, salario := range salarios {
		fmt.Printf("\nO salário de %s é de: R$%d", nome, salario)
	}
}
