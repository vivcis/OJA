package main

import (
	"fmt"
	"github.com/decadevs/shoparena/server"
)

func main() {
	fmt.Println("Starting... Oja")

	err := server.Start()
	if err != nil {
		fmt.Println("Error starting server in main", err)
		return
	}

}
