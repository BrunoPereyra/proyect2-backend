package Post

import (
	"backend/database"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPost(c *fiber.Ctx) error {

	// connect database
	db, err := database.GoMongoDB()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	// pipeline de agregación
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "userid",
				"foreignField": "_id",
				"as":           "User",
			},
		},
		{
			"$project": bson.M{
				"_id":           1,
				"status":        1,
				"postimage":     1,
				"Likes":         1,
				"timestamp":     1,
				"User._id":      1,
				"User.nameuser": 1,
				"User.email":    1,
				"User.avatar":   1,
			},
		},
	}

	// ejecutar agregación
	cursor, err := db.Collection("post").Aggregate(context.Background(), pipeline)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	// iterar el cursor para obtener los documentos
	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	// Devolver la respuesta de la funcion
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    results,
	})
}
