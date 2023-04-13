package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/jwt"
	"backend/models"
	"backend/validator"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c *fiber.Ctx) error {

	var DataForLogin validator.LoginValidatorStruct

	if err := c.BodyParser(&DataForLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	if err := DataForLogin.LoginValidator(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	Database, errDB := database.GoMongoDB()
	if errDB != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	// user exist?
	GoMongoDBCollUsers := Database.Collection("users")
	findUserLogin := bson.D{
		{Key: "nameuser", Value: DataForLogin.NameUser},
	}
	var result models.UserModel
	err := GoMongoDBCollUsers.FindOne(context.TODO(), findUserLogin).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
	}
	//password incorrect
	if err := helpers.DecodePassword(result.PasswordHash, DataForLogin.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	token, err := jwt.CreateToken(result)
	if err != nil {
		log.Fatal("Login,CreateTokenError", err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Token created",
		"token":   token,
	})
}
