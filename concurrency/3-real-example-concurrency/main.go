package main

import (
	"fmt"
	"net/http"
	"time"
)

var numbervisits uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numbervisits++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte("Essa página foi visitada " + string(fmt.Sprintf("%d", numbervisits)) + " vezes"))
	})

	http.ListenAndServe(":3000", nil)

}
