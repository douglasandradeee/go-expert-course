package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)

}

func publish(ch chan int) {
	// escreve no canal
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int) {
	// faz a leitura do canal e atribui o valor a variavel x
	for x := range ch {
		fmt.Println("valores recebidos do publish: ", x)
	}
}
