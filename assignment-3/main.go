package main

import (
	"assignment-3/controllers"
	"assignment-3/routes"
	"log"
)

func main() {
	server := routes.InitServer(&controllers.Controllers{})
	err := server.Run(":8080")
	if err != nil {
		log.Println("failed to start server")
	}
}
