package main

import (
	"fmt"

	"github.com/douglasandradeee/go-expert-course/apis/configs"
)

func main() {

	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
