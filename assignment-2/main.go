package main

import (
	"assignment-2/controllers"
	"assignment-2/database"
	"assignment-2/routes"
	"log"
)

func main() {
	db := database.InitDb()

	c := &controllers.Controllers{Database: db}

	server := routes.InitServer(c)
	err := server.Run(":8080")
	if err != nil {
		log.Println("failed to start server")
	}
}
