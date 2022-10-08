package main

import (
	"assignment-3/controllers"
	"assignment-3/models"
	"assignment-3/repositories"
	"assignment-3/routes"
	"log"
	"math/rand"
	"time"
)

func main() {
	go updateStatus()

	server := routes.InitServer(&controllers.Controllers{})
	err := server.Run(":8080")
	if err != nil {
		log.Println("failed to start server")
	}
}

func updateStatus() {
	for {
		status := models.Status{
			Status: models.Weather{
				Water: rand.Float64() * 100,
				Wind:  rand.Float64() * 100,
			},
		}
		repositories.SetStatus(status)

		time.Sleep(15 * time.Second)
	}
}
