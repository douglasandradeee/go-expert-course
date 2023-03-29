package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	c := http.Client{
		Timeout: time.Second,
	}

	resp, err := c.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", string(bodyByte))
}
