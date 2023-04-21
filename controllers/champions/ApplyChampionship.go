package champions

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChampionshipId struct {
	Championship_id string `json:"Championship_id"`
}

func ApplyChampionship(c *fiber.Ctx) error {

	var req ChampionshipId
	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}
	db, errdb := database.GoMongoDB()
	if errdb != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "StatusInternalServerError",
		})
	}

	// Parsear la solicitud de aplicación de JSON
	dataMiddleware := c.Context().UserValue("nameUser")
	dataMiddlewareString, _ := dataMiddleware.(string)
	user, err := helpers.UserTMiddlExist(dataMiddlewareString, db)

	if err != nil {
		return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	id, errorID := primitive.ObjectIDFromHex(req.Championship_id)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id unrecognized",
		})
	}

	var Championship models.Championships
	findchampion := bson.D{
		{Key: "_id", Value: id},
	}
	CollectionChampionship := db.Collection("championship")
	errFindChampionship := CollectionChampionship.FindOne(context.TODO(), findchampion).Decode(&Championship)

	if errFindChampionship != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Championship not found"})
	}

	for i, ApplicantsId := range Championship.Applicants {
		fmt.Println(i)
		if ApplicantsId == user.ID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Ya has aplicado a este campeonato",
			})

		}
	}

	Championship.Applicants = append(Championship.Applicants, user.ID)
	update := bson.M{
		"$set": bson.M{
			"applicants": Championship.Applicants,
		},
	}

	_, err = CollectionChampionship.UpdateOne(context.TODO(), findchampion, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{
		"message": "Participant added successfully",
	})

}
