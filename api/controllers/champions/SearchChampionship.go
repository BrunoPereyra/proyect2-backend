package champions

import (
	"backend/api/models"
	"backend/database"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchChampions(c *fiber.Ctx) error {
	refChampionship := c.Query("refChampionship")
	if len(refChampionship) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}
	db, err := database.GoMongoDB()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	championshipCollecion := db.Collection("championship")
	regex := primitive.Regex{Pattern: refChampionship, Options: "i"}
	findChampionships := bson.D{
		{Key: "name", Value: regex},
	}
	cursor, err := championshipCollecion.Find(context.TODO(), findChampionships)
	if err != nil {
		c.JSON(fiber.Map{
			"message": "NotFound championship",
		})
	}
	var championships []models.Championships
	for cursor.Next(context.TODO()) {
		var championship models.Championships
		err := cursor.Decode(&championship)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error server",
			})
		}
		championships = append(championships, championship)
	}
	if len(championships) <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "NotFound",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    championships,
	})
}
