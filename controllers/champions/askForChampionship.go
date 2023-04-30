package champions

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Idchampionship struct {
	IDchampionship string `json:"IDchampionship"`
}

func AskForChampionship(c *fiber.Ctx) error {
	// connect database
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	var req Idchampionship
	c.BodyParser(&req)

	collectionchampionship := databaseGoMongodb.Collection("championship")
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
