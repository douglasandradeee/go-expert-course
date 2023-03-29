package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	c := http.Client{}

	req, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyBytes))

}
