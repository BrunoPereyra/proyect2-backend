package controllers

import (
	"backend/database"
	"backend/models"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SearchUserData struct {
	Nameuser string `json:"nameuser"`
}

func SearchUser(c *fiber.Ctx) error {
	var SearchUserData SearchUserData
	fmt.Println(SearchUserData, "--")
	if err := c.BodyParser(&SearchUserData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())

	databaseGoMongodb := db.Pool.Database("goMoongodb")

	collecionusers := databaseGoMongodb.Collection("users")
	options := options.Find().
		SetLimit(4)
	regex := primitive.Regex{Pattern: SearchUserData.Nameuser, Options: "i"}
	findnameuser := bson.D{
		{Key: "nameuser", Value: regex},
	}
	cursor, errfinduser := collecionusers.Find(context.TODO(), findnameuser, options)
	if errfinduser != nil {
		c.JSON(fiber.Map{
			"message": "NotFound championship",
		})
	}
	var users []models.User
	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error server",
			})
		}
		users = append(users, user)
	}
	if len(users) <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "NotFound",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    users,
	})
}
