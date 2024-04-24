package main

import "fmt"

//Thread 1, thread principal
func main() {
	channel := make(chan string) //canal vazio

	//Thread 2
	go func() {
		// canal cheio e ou sendo preenchido
		//preenchendo o canal
		channel <- "Hello World"
	}()

	msg := <-channel // canal esvaziando
	fmt.Println(msg)

}
