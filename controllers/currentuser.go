package controllers

import (
	"backend/database"
	"backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Currentuser(c *fiber.Ctx) error {
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	nameUser := c.Context().UserValue("nameUser")
	GoMongoDBCollUsers := databaseGoMongodb.Collection("users")
	findUser := bson.D{
		{Key: "nameuser", Value: nameUser},
	}
	var user models.User
	errGoMongoDBCollUsers := GoMongoDBCollUsers.FindOne(context.TODO(), findUser).Decode(&user)
	if errGoMongoDBCollUsers != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "StatusInternalServerError",
		})
	}
	if user.NameUser == nameUser {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "StatusOK",
			"data":    user,
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "NotFoundUser",
		})
	}
}
