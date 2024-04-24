package main

import "fmt"

func main() {

	//Canal com buffer
	chWithBuffer := make(chan string, 2)
	chWithBuffer <- "hello"
	chWithBuffer <- "world"

	fmt.Println(<-chWithBuffer)
	fmt.Println(<-chWithBuffer)

	//Simulação de erro deadlock
	// chWithoutBuffer := make(chan string)
	// chWithoutBuffer <- "hello"
	// chWithoutBuffer <- "world"

	// fmt.Println(<-chWithoutBuffer)
	// fmt.Println(<-chWithoutBuffer)
}
