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

type ReqData struct {
	UserID         string `json:"user_id"`
	ChampionshipID string `json:"Championship_id"`
}

func ParticipantsAwaitingForPayment(c *fiber.Ctx) error {
	var req ReqData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad Request"})
	}

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

	idChampionship, errorID := primitive.ObjectIDFromHex(req.ChampionshipID)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id StatusBadRequest",
		})
	}
	// Championship existe?
	var Championship models.Championships
	findchampion := bson.D{
		{Key: "_id", Value: idChampionship},
	}
	CollectionChampionship := databaseGoMongodb.Collection("championship")
	errFindChampionship := CollectionChampionship.FindOne(context.TODO(), findchampion).Decode(&Championship)
	if errFindChampionship != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Championship not found"})
	}

	// id del usuario de string a objectId
	UserIDReq, errorID := primitive.ObjectIDFromHex(req.UserID)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id StatusBadRequest",
		})
	}
	// el usuario del token es el dueño del Championship?
	var user models.User
	select {
	case user = <-UserCreator:
		if Championship.Creator != user.ID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden"})
		}
	case <-errChanelUserTMiddlExist:
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "StatusNotAcceptable",
		})
	}

	for _, value := range Championship.AcceptedApplicants {
		if value == UserIDReq {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "is already participating",
			})
		}
	}

	// actualizar evento

	Championship.AcceptedApplicants = append(Championship.AcceptedApplicants, UserIDReq)
	update := bson.M{
		"$pull": bson.M{
			"applicants": UserIDReq,
		},
		"$set": bson.M{
			"acceptedApplicants": Championship.AcceptedApplicants,
		},
	}

	_, errUpdateOne := CollectionChampionship.UpdateOne(context.TODO(), findchampion, update)
	if errUpdateOne != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Participant added successfully",
	})

}
