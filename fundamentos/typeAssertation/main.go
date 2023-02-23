package main

import (
	"log"
)

func main() {
	var minhaVar interface{} = "Douglas Andrade"
	println(minhaVar.(string))
	res, ok := minhaVar.(string)
	if !ok {
		log.Fatalf("O valor de res é: %v, e o resultado de ok é: %v", res, ok)
	}
	//fmt.Printf("O valor de res é: %v, e o resultado de ok é: %v", res, ok)
	//fmt.Println(minhaVar.(string))
}
