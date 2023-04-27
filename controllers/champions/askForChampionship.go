package champions

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Idchampionship struct {
	IDchampionship string `json:"IDchampionship"`
}

func AskForChampionship(c *fiber.Ctx) error {
	// connect database
	db, err := database.GoMongoDB()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	var req Idchampionship
	c.BodyParser(&req)

	collectionchampionship := db.Collection("championship")
	id, errorID := primitive.ObjectIDFromHex(req.IDchampionship)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}

	find := bson.D{
		{Key: "_id", Value: id},
	}
	var collectionChampionship models.Championships
	errfindChampion := collectionchampionship.FindOne(context.TODO(), find).Decode(&collectionChampionship)

	if errfindChampion != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	fmt.Println(collectionChampionship)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    collectionChampionship,
	})
}
