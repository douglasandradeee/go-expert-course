package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}

	bodyBytesBuffer := bytes.NewBuffer([]byte(`Name:"Douglas", Age:31`))
	resp, err := c.Post("https://www.google.com.br", "application/json", bodyBytesBuffer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// OU
	// faz uma copia da response no terminal, poderia fazer isso numa pagina de nosso servidor web.
	io.CopyBuffer(os.Stdout, resp.Body, nil)

	fmt.Printf("%+v", string(bodyByte))
}
