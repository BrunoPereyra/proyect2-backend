package championshipAdmin

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"log"

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
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	// user existe token?
	dataMiddleware := c.Context().UserValue("nameUser")
	UserCreator := make(chan models.User)
	errChanelUserTMiddlExist := make(chan error)
	go helpers.UserTMiddlExist(dataMiddleware.(string), databaseGoMongodb, UserCreator, errChanelUserTMiddlExist)

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
	CollectionChampionship := databaseGoMongodb.Collection("championship")
	errFindChampionship := CollectionChampionship.FindOne(context.TODO(), findchampion).Decode(&Championship)

	if errFindChampionship != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Championship not found",
		})
	}
	// existe en Applicants
	var user models.User
	select {
	case user = <-UserCreator:
		for _, ApplicantsId := range Championship.Applicants {
			if ApplicantsId == user.ID {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Ya has aplicado a este campeonato",
				})
			}
		}
	case <-errChanelUserTMiddlExist:
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "StatusNotAcceptable",
		})
	}
	// existe en Participants
	for _, ApplicantsId := range Championship.AcceptedApplicants {
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "Participant added successfully",
		})
	}
}
