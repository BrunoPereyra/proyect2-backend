package controllers

import (
	"backend/api/models"
	"backend/database"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dataGetChampionshipStruct struct {
	Page int `json:"page"`
}

func GetChampionship(c *fiber.Ctx) error {

	PageSize := 10

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
	skip := int64((dataGetChampionship.Page - 1) * PageSize)

	// Calcular el número de documentos que se deben devolver
	limit := int64(PageSize)

	// Crear las opciones de búsqueda
	options := options.Find().
		SetSkip(skip).
		SetLimit(limit)

	// Obtener el cursor de los documentos que cumplen con los criterios de búsqueda
	cursor, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	var championships []models.Championships
	// Iterar el cursor para obtener los documentos paginados
	for cursor.Next(context.Background()) {
		var championship models.Championships
		err := cursor.Decode(&championship)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "StatusServiceUnavailable",
			})
		}
		championships = append(championships, championship)
	}

	// Devolver la respuesta
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    championships,
	})
}
