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

func CreateEvent(c *fiber.Ctx) error {
	dataMiddleware := c.Context().UserValue("nameUser")
	client := database.DB()

	db := client.Database("goMoongodb").Collection("users")

	// usuario del middleware existe?
	find := bson.D{
		{Key: "nameuser", Value: dataMiddleware},
	}
	var UserCreator models.UserModel
	err := db.FindOne(context.TODO(), find).Decode(&UserCreator)
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
	var validateEventData validator.Event
	if err := c.BodyParser(&validateEventData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}
	// validateEventData.Creator = UserCreator.ID
	if err := validateEventData.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	var modelEvent models.Event

	modelEvent.Creator = validateEventData.Creator
	modelEvent.Description = validateEventData.Description
	modelEvent.Name = validateEventData.Name
	modelEvent.Prize = validateEventData.Prize
	modelEvent.Entry = validateEventData.Entry
	modelEvent.Requirements = validateEventData.Requirements

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "error server",
	})
}
