package main

import (
	"fmt"

	"github.com/jmdalmeida/caracolines"
)

func main() {
	//config, _ := caracolines.LoadConfig("./test/config.yaml")
	services, _ := caracolines.LoadServices("./test/services.yaml")

	for _, item := range services.Items {
		fmt.Println(item)
	}
}
