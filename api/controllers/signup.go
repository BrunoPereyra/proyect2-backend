package controllers

import (
	"backend/api/helpers"
	"backend/api/models"
	"backend/api/validator"
	"backend/database"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Signup(c *fiber.Ctx) error {

	var newUser validator.UserModelValidator
	if err := c.BodyParser(&newUser); err != nil {
		return c.JSON(fiber.Map{
			"res": "error a cargar los datos",
		})
	}
	if err := newUser.ValidateUserFind(); err != nil {
		return c.JSON(fiber.Map{
			"res": "malformed request",
			"err": err,
		})
	}
	client := database.DB()
	db := client.Database("goMoongodb").Collection("users")

	findUserInDb := bson.D{
		{Key: "NameUser", Value: newUser.NameUser},
		{Key: "Email", Value: newUser.Email},
	}

	if resFind, err := db.Find(context.TODO(), findUserInDb, options.Find().SetLimit(1)); err != nil {
		if err == mongo.ErrNoDocuments {

			passwordHash, err := helpers.HashPassword(newUser.Password)

			if err != nil {
				return c.JSON(fiber.Map{
					"res": "server error",
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

			user, err := db.InsertOne(context.TODO(), modelNewUser)
			if err != nil {
				return c.JSON(fiber.Map{
					"res": "error save user",
				})
			}
			return c.JSON(fiber.Map{
				"res": user,
			})

		} else {
			return c.JSON(fiber.Map{
				"res": "server error",
			})
		}
	} else {
		fmt.Println(resFind)
		return c.JSON(fiber.Map{
			"res": "exist NameUser or Email",
		})
	}

}
