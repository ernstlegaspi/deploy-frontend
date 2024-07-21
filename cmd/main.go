package main

import (
	"fmt"
	"test/cmd/api"
)

func main() {
	server := api.CreateServer()

	if err := server.InitAPI(":3001"); err != nil {
		fmt.Println("Erororro")
		panic(err)
	}
}
