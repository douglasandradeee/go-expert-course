package main

import "fmt"

func main() {
	hello := make(chan string)
	go receive("hello", hello)
	read(hello)
}

// essa função tem um canal direcional que apenas recebe informações
func receive(name string, hello chan<- string) {
	hello <- name
}

// essa função tem um canal direcional que apenas ler/entrega informações
func read(data <-chan string) {
	fmt.Println(<-data)
}
