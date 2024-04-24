package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		// toda vez que passarmos pelo done iremos decrementar a quantidade de operações/tarefas a serem executadas no agrupamento
		wg.Done()
	}
}

// a função main é a goroutine principal ou seja é minha primeira thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	//adiciona qtidade de operações/tarefas a serem executadas no agrupamento
	waitGroup.Add(25)
	//thread 2
	go task("A", &waitGroup)
	//thread 3
	go task("B", &waitGroup)
	// thread 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	// enquanto as operações setadas no agrupamento não forem executadas o programa não termina, iremos esperar.
	waitGroup.Wait()
}
