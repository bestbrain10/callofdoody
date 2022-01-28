package handlers

import (
	"gofiberme/config"
	"gofiberme/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCars(c *fiber.Ctx) error {
	query := bson.D{}
	data, err := config.Collections.Cars.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var allDogs []entities.ICars = make([]entities.ICars, 0)

	if err := data.All(c.Context(), &allDogs); err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(&allDogs)
}
