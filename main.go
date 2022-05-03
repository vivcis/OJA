package main

import (
	"fmt"

	"github.com/decadevs/shoparena/database"
	"github.com/decadevs/shoparena/router"
)

func main() {
	fmt.Println("Starting... Oja")

	values := database.InitDBParams()

	var PDB = new(database.PostgresDb)
	PDB.Init(values.Host, values.User, values.Password, values.DbName, values.Port)
	router, port := router.SetupRouter()
	fmt.Println("connected on port ", port)
	router.Run(port)

}
