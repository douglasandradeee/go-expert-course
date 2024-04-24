package main

import (
	"log"
	"net/http"
)

// File Server - É um servidor de arquivos estáticos.
func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello form blog"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
