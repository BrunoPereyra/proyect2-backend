package controllers

import (
	"backend/database"
	"backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Currentuser(c *fiber.Ctx) error {
	Database, errDB := database.GoMongoDB()
	nameUser := c.Context().UserValue("nameUser")
	if errDB != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	GoMongoDBCollUsers := Database.Collection("users")
	findUser := bson.D{
		{Key: "nameuser", Value: nameUser},
	}
	var user models.UserModel
	err := GoMongoDBCollUsers.FindOne(context.TODO(), findUser).Decode(&user)
	if err != nil {
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
