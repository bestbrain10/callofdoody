package handlers

import (
	"gofiberme/config"
	"gofiberme/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// You can godoc functions
// Gets a single car, returns error if any
func GetCar(c *fiber.Ctx) error {
	carId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	var singleCar bson.M
	query := bson.D{{"_id", carId}}
	if err := config.Collections.Cars.FindOne(c.Context(), query).Decode(&singleCar); err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(singleCar)
}

func AddCar(c *fiber.Ctx) error {
	car := new(entities.ICars)
	c.BodyParser(car)
	response, err := config.Collections.Cars.InsertOne(c.Context(), car)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	var insertedCar bson.M
	query := bson.D{{"_id", response.InsertedID}}
	if err := config.Collections.Cars.FindOne(c.Context(), query).Decode(&insertedCar); err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(insertedCar)
}

func UpdateCar(c *fiber.Ctx) error {
	carId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	car := new(entities.ICars)
	if err := c.BodyParser(car); err != nil {
		return c.Status(500).JSON(err)
	}

	query := bson.D{{"_id", carId}}
	body := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{"name", car.Name},
				{"model", car.Model},
				{"brand", car.Brand},
				{"isGood", car.IsGood},
				{"mileage", car.Mileage},
			},
		},
	}
	var updatedCar bson.M

	if err := config.Collections.Cars.FindOneAndUpdate(c.Context(), &query, &body).Decode(&updatedCar); err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(updatedCar)
}

func RemoveCar(c *fiber.Ctx) error {
	carId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	query := bson.D{{"_id", carId}}
	var car bson.M
	if err := config.Collections.Cars.FindOneAndDelete(c.Context(), &query).Decode(&car); err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(car)
}
