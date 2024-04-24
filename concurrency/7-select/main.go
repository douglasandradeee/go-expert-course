package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	//rabbitmq
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{i, "hello from rabbitmq"}
			c1 <- msg
		}
	}()

	//Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{i, "hello from Kafka"}

			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println("received from RabbitMQ ID:", msg.Id, "-", msg.Msg)
		case msg := <-c2:
			fmt.Println("received from Kafka ID:", msg.Id, "-", msg.Msg)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout")
			// default:
			// 	fmt.Println("no activity")
		}
	}

}
