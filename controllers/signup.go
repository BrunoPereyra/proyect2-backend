package controllers

import (
	"backend/config"
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/validator"
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Signup(c *fiber.Ctx) error {

	var newUser validator.UserModelValidator
	fileHeader, errfileHeader := c.FormFile("avatar")
	if errfileHeader != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}
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

	Database, errDB := database.GoMongoDB()
	if errDB != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	// password
	passwordHashChan := make(chan string)
	go helpers.HashPassword(newUser.Password, passwordHashChan)

	// buscar usuario por name user or gmail
	GoMongoDBCollUsers := Database.Collection("users")
	findUserInDb := bson.D{
		{
			Key: "$or",
			Value: bson.A{
				bson.D{{Key: "nameuser", Value: newUser.NameUser}},
				bson.D{{Key: "email", Value: newUser.Email}},
			},
		},
	}
	var findUserInDbExist models.User
	err := GoMongoDBCollUsers.FindOne(context.TODO(), findUserInDb).Decode(&findUserInDbExist)

	if err != nil {
		// si no exiaste crealo
		if err == mongo.ErrNoDocuments {

			passwordHash := <-passwordHashChan
			if passwordHash == "error" {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error hash",
				})
			}
			file, _ := fileHeader.Open()

			ctx := context.Background()
			cldService, _ := cloudinary.NewFromURL(config.CLOUDINARY_URL())
			resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
			if passwordHash == "error" {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error hash",
				})
			}
			avatarUrl := resp.SecureURL
			var modelNewUser models.User
			fmt.Println(avatarUrl)

			modelNewUser.Avatar = avatarUrl
			modelNewUser.FullName = newUser.FullName
			modelNewUser.NameUser = newUser.NameUser
			modelNewUser.PasswordHash = passwordHash
			modelNewUser.Pais = newUser.Pais
			modelNewUser.Ciudad = newUser.Ciudad
			modelNewUser.Email = newUser.Email
			modelNewUser.Instagram = newUser.Instagram
			modelNewUser.Twitter = newUser.Twitter
			modelNewUser.Youtube = newUser.Youtube
			// incertar usuario
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
			// server error
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
