package main

import (
	"fmt"
	"github.com/decadevs/shoparena/database"
	"github.com/decadevs/shoparena/router"
)

func main() {
	fmt.Println("Oja")

	var PDB = new(database.PostgresDb)
	PDB.Init()
	router, port := router.SetupRouter()
	fmt.Println("connected on port ", port)
	router.Run(port)

}
