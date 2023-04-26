package champions

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/validator"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateChampionship(c *fiber.Ctx) error {

	Database, errDB := database.GoMongoDB()
	if errDB != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}

	// usuario del middleware existe?
	dataMiddleware := c.Context().UserValue("nameUser")
	dataMiddlewareString, _ := dataMiddleware.(string)

	UserCreator, err := helpers.UserTMiddlExist(dataMiddlewareString, Database)
	if err != nil {
		return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// creacion de championshipsValidate
	var championshipsValidate validator.ChampionshipsValidate
	if err := c.BodyParser(&championshipsValidate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
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
	modelChampionships.CreatedAt = time.Now()
	modelChampionships.UpdatedAt = time.Now()
	modelChampionships.Applicants = []primitive.ObjectID{}
	modelChampionships.Participants = []primitive.ObjectID{}
	modelChampionships.Votesoftheparticipants = make(map[primitive.ObjectID][]primitive.ObjectID)
	modelChampionships.Voters = []primitive.ObjectID{}

	Championshipdb := Database.Collection("championship")

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
