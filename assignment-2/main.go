package main

import (
	"assignment-2/controllers"
	"assignment-2/database"
	"assignment-2/routes"
)

func main() {
	database := database.InitDb()

	controllers := &controllers.Controllers{Database: database}

	server := routes.InitServer(controllers)
	server.Run(":8080")
}
