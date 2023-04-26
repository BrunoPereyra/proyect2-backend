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

type Vote struct {
	Championship string `json:"championship"`
	IDUser       string `json:"iduser"`
}

func VoteForChampionship(c *fiber.Ctx) error {
	var vote Vote
	if errBodyParser := c.BodyParser(&vote); errBodyParser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}
	ParticipantTheUserVotesFor, iduserObjectIDFromHex := primitive.ObjectIDFromHex(vote.IDUser)
	if iduserObjectIDFromHex != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id unrecognized",
		})
	}
	db, errDB := database.GoMongoDB()
	if errDB != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}
	collecionChampionship := db.Collection("championship")

	ChampionshipId, idChampionshipObjectIDFromHex := primitive.ObjectIDFromHex(vote.Championship)
	if idChampionshipObjectIDFromHex != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id unrecognized",
		})
	}

	findCollection := bson.D{
		{Key: "_id", Value: ChampionshipId},
	}
	var Championship models.Championships
	errcollecionChampionship := collecionChampionship.FindOne(context.TODO(), findCollection).Decode(&Championship)
	if errcollecionChampionship != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Championship not found",
		})
	}

	dataMiddleware := c.Context().UserValue("nameUser")
	dataMiddlewareString, _ := dataMiddleware.(string)
	user, errhelpers := helpers.UserTMiddlExist(dataMiddlewareString, db)

	if errhelpers != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id unrecognized",
		})
	}

	for _, voter := range Championship.Voters {
		if voter == user.ID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "you already voted",
			})
		}
	}
	// si el existe el usuario por el que vota en Votesoftheparticipants, ok = true
	ParticipantTheUserVotesForOk, ok := Championship.Votesoftheparticipants[ParticipantTheUserVotesFor]
	if ok {
		ParticipantTheUserVotesForOk = append(ParticipantTheUserVotesForOk, user.ID)
		Championship.Votesoftheparticipants[ParticipantTheUserVotesFor] = ParticipantTheUserVotesForOk

	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "el participante al que queres votar no existe",
		})
	}
	Championship.Voters = append(Championship.Voters, user.ID)
	update := bson.M{
		"$set": bson.M{
			"Voters":                 Championship.Voters,
			"Votesoftheparticipants": Championship.Votesoftheparticipants,
		},
	}

	_, errCollecionChampionship := collecionChampionship.UpdateOne(context.TODO(), findCollection, update)
	if errCollecionChampionship != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	}
}
