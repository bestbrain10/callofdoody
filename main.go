package main

import (
	"context"
	"gofiberme/config"
	"gofiberme/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if err := config.ConnectDatabase(); err != nil {
		panic(err)
	}

	defer config.Client.Disconnect(context.Background())
	app.Get("/cars", handlers.GetCars)
	app.Get("/cars/:id", handlers.GetCar)
	app.Delete("/cars/:id", handlers.RemoveCar)
	app.Post("/cars", handlers.AddCar)
	app.Put("/cars/:id", handlers.UpdateCar)

	app.Listen(":6700")
}
