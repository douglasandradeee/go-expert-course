package main

import "fmt"

/*
Uma interface vazia significa que podemos ter um pouco mais de liberadade nos tipos
de dados que queremos passar. E também significa que essa interface implementará todo mundo.

CUIDADO E MODERAÇÃO ao utilizar desse artificio, pois o go é uma linguagem fortemente tipada.
Esse artificio serve pra quebrar isso de certa forma.
*/
type x interface{}

func main() {

	var z interface{} = 10
	var y interface{} = "Hello World"
	showType(z)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é: %T, e o valor é: %v\n", t, t)
}
