package controllers

import (
	"backend/api/models"
	"backend/api/validator"
	"backend/database"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChampionship(c *fiber.Ctx) error {
	dataMiddleware := c.Context().UserValue("nameUser")
	Database := database.GoMongoDB()

	GoMongoDBCollUsers := Database.Collection("users")

	// usuario del middleware existe?
	find := bson.D{
		{Key: "nameuser", Value: dataMiddleware},
	}
	var UserCreator models.UserModel
	err := GoMongoDBCollUsers.FindOne(context.TODO(), find).Decode(&UserCreator)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
				"message": "user not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error server",
			})
		}
	}
	// ---------- 0 -------------
	var championshipsValidate validator.ChampionshipsValidate
	if err := c.BodyParser(&championshipsValidate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}

	championshipsValidate.Creator = UserCreator.ID
	if err := championshipsValidate.ChampionshipsValidate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	var modelChampionships models.Championships

	modelChampionships.Creator = championshipsValidate.Creator
	modelChampionships.Description = championshipsValidate.Description
	modelChampionships.Name = championshipsValidate.Name
	modelChampionships.Prize = championshipsValidate.Prize
	modelChampionships.Entry = championshipsValidate.Entry
	modelChampionships.Requirements = championshipsValidate.Requirements

	Championshipdb := Database.Collection("Championship")

	event, err := Championshipdb.InsertOne(context.TODO(), modelChampionships)
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
			"data":    event,
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "error server",
		"err":     err,
	})
}
