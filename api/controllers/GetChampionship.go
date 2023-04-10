package controllers

import (
	"backend/api/models"
	"backend/database"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dataGetChampionshipStruct struct {
	Page int `json:"page"`
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
	// Obtener el número total de documentos
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	// buscar id
	var IDCreatedAt models.Championships

	id, _ := primitive.ObjectIDFromHex("643442b8ca28664105e895ed")
	errIDCreatedAt := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&IDCreatedAt)
	if errIDCreatedAt != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	// buscar 10 adelante de id
	fmt.Println(IDCreatedAt.CreatedAt)
	findNextTheId := bson.M{
		"createdat": bson.M{"$gt": IDCreatedAt.CreatedAt},
	}
	options := options.Find().
		SetSort(bson.M{"createdat": 1}).
		SetLimit(3)

	cursor, err := collection.Find(context.TODO(), findNextTheId, options)

	// Iterar el cursor para obtener los documentos paginados
	var championships []models.Championships
	for cursor.Next(context.TODO()) {
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
