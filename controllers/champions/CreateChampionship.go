package champions

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/validator"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateChampionship(c *fiber.Ctx) error {

	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	// usuario del middleware existe?
	dataMiddleware := c.Context().UserValue("nameUser")
	UserCreatorChan := make(chan models.User)
	errChanelUserTMiddlExist := make(chan error)
	go helpers.UserTMiddlExist(dataMiddleware.(string), databaseGoMongodb, UserCreatorChan, errChanelUserTMiddlExist)

	// creacion de championshipsValidate
	var championshipsValidate validator.ChampionshipsValidate
	if err := c.BodyParser(&championshipsValidate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	var UserCreator models.User
	select {
	case UserCreator = <-UserCreatorChan:
		championshipsValidate.Creator = UserCreator.ID
		if err := championshipsValidate.ChampionshipsValidate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
				"error":   err.Error(),
			})
		}
	case <-errChanelUserTMiddlExist:
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "StatusNotAcceptable",
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
	modelChampionships.AcceptedApplicants = []primitive.ObjectID{}

	modelChampionships.ParticipantsWhoPaidTheEntrance = []models.ParticipantsWhoPaidTheEntrance{}

	modelChampionships.Votesoftheparticipants = make(map[primitive.ObjectID][]primitive.ObjectID)
	modelChampionships.Voters = []primitive.ObjectID{}

	Championshipdb := databaseGoMongodb.Collection("championship")

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
