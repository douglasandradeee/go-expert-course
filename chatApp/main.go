package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type MessageDTO struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var broadcast = make(chan MessageDTO)
var Clients = make(map[*websocket.Conn]bool)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", HandleConnections)

	go HandleMessages()

	fmt.Println("Server running")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer ws.Close()

	Clients[ws] = true

	for {
		var msg MessageDTO
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(Clients, ws)
			break
		}
		broadcast <- msg
	}

}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				//client.Close()
				delete(Clients, client)
				client.Close()
			}
		}
	}
}
