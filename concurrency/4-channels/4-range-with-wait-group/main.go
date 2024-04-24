package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg)

	wg.Wait()

}

func publish(ch chan int) {
	// escreve no canal
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	// faz a leitura do canal e atribui o valor a variavel x
	for x := range ch {
		fmt.Println("valores recebidos do publish: ", x)
		wg.Done()
	}
}
