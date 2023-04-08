package controllers

import (
	"backend/api/helpers"
	"backend/api/models"
	"backend/api/validator"
	"backend/database"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(c *fiber.Ctx) error {

	var newUser validator.UserModelValidator
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}

	if err := newUser.ValidateUserFind(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	GoMongoDB := database.GoMongoDB()
	GoMongoDBCollUsers := GoMongoDB.Collection("users")

	findUserInDb := bson.D{
		{
			Key: "$or",
			Value: bson.A{
				bson.D{{Key: "nameuser", Value: newUser.NameUser}},
				bson.D{{Key: "email", Value: newUser.Email}},
			},
		},
	}
	var findUserInDbExist models.UserModel
	err := GoMongoDBCollUsers.FindOne(context.TODO(), findUserInDb).Decode(&findUserInDbExist)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			passwordHash, err := helpers.HashPassword(newUser.Password)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
					"err":     err,
				})
			}

			var modelNewUser models.UserModel

			modelNewUser.Avatar = newUser.Avatar
			modelNewUser.FullName = newUser.FullName
			modelNewUser.NameUser = newUser.NameUser
			modelNewUser.PasswordHash = passwordHash
			modelNewUser.Pais = newUser.Pais
			modelNewUser.Ciudad = newUser.Ciudad
			modelNewUser.Email = newUser.Email
			modelNewUser.Instagram = newUser.Instagram
			modelNewUser.Twitter = newUser.Twitter
			modelNewUser.Youtube = newUser.Youtube

			user, err := GoMongoDBCollUsers.InsertOne(context.TODO(), modelNewUser)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
					"err":     err,
				})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": user,
			})

		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
				"err":     err,
			})
		}
	} else {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "exist NameUser or Email",
		})
	}

}
