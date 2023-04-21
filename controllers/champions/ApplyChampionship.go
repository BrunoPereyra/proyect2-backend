package champions

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChampionshipId struct {
	ChampionshipID string `json:"Championship_id"`
}

func ApplyChampionship(c *fiber.Ctx) error {
	var req ChampionshipId
	err := c.BodyParser(&req)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}
	// conexión db
	db, errdb := database.GoMongoDB()
	if errdb != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "StatusInternalServerError",
		})
	}

	// user existe token?
	dataMiddleware := c.Context().UserValue("nameUser")
	dataMiddlewareString, _ := dataMiddleware.(string)
	user, err := helpers.UserTMiddlExist(dataMiddlewareString, db)

	if err != nil {
		return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	// ChampionshipID de string a objectId
	id, errorID := primitive.ObjectIDFromHex(req.ChampionshipID)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id unrecognized",
		})
	}
	// Championship al que aplica
	var Championship models.Championships
	findchampion := bson.D{
		{Key: "_id", Value: id},
	}
	CollectionChampionship := db.Collection("championship")
	errFindChampionship := CollectionChampionship.FindOne(context.TODO(), findchampion).Decode(&Championship)

	if errFindChampionship != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Championship not found"})
	}
	// existe en Applicants
	for _, ApplicantsId := range Championship.Applicants {
		if ApplicantsId == user.ID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Ya has aplicado a este campeonato",
			})

		}
	}
	// existe en Participants
	for _, ApplicantsId := range Championship.Participants {
		if ApplicantsId == user.ID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Ya estas participando en el campeonato",
			})

		}
	}
	// update del doumento
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
