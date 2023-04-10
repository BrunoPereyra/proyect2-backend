package controllers

import (
	"backend/api/models"
	"backend/database"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dataGetChampionshipStruct struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func GetChampionship(c *fiber.Ctx) error {
	var dataGetChampionship dataGetChampionshipStruct
	if err := c.BodyParser(&dataGetChampionship); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	db, err := database.GoMongoDB()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	collection := db.Collection("championship")

	// Calcular la cantidad de documentos que se deben omitir
	skip := int64((dataGetChampionship.Page - 1) * dataGetChampionship.PageSize)

	// Calcular el número de documentos que se deben devolver
	limit := int64(dataGetChampionship.PageSize)

	// Crear las opciones de búsqueda
	options := options.Find().
		SetSkip(skip).
		SetLimit(limit)

	cursor, err := collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	var championships []models.Championships
	if err := cursor.All(context.Background(), &championships); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	fmt.Println(len(championships))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    championships,
	})
}
