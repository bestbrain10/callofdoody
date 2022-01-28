package main

import (
	"context"
	"gofiberme/config"
	"gofiberme/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	app := fiber.New()
	dbURL := goDotEnvVariable("DB_URL")
	if err := config.ConnectDatabase(dbURL); err != nil {
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
