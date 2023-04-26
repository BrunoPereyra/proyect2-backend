package championshipAdmin

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReqData struct {
	UserID         string `json:"user_id"`
	ChampionshipID string `json:"Championship_id"`
}

func BecomeParticipant(c *fiber.Ctx) error {
	var req ReqData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad Request"})
	}

	db, err := database.GoMongoDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	// user existe token?
	dataMiddleware := c.Context().UserValue("nameUser")
	dataMiddlewareString, _ := dataMiddleware.(string)
	user, errUserTMiddlExist := helpers.UserTMiddlExist(dataMiddlewareString, db)
	if errUserTMiddlExist != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	// ChampionshipID de string a objectId
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
	CollectionChampionship := db.Collection("championship")
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
	if Championship.Creator != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden"})
	}
	// inicializar Votesoftheparticipants
	if Championship.Votesoftheparticipants == nil {
		Championship.Votesoftheparticipants = make(map[primitive.ObjectID][]primitive.ObjectID)
	}

	// si Championship.Votesoftheparticipants[user.ID] no existe
	if _, ok := Championship.Votesoftheparticipants[user.ID]; !ok {
		// actualizar evento
		update := bson.M{
			"$pull": bson.M{
				"applicants": UserIDReq,
			},
			"$addToSet": bson.M{
				"participants": UserIDReq,
			},
			"$set": bson.M{
				"Votesoftheparticipants": bson.M{
					user.ID.Hex(): []primitive.ObjectID{},
				},
			},
		}

		_, errUpdateOne := CollectionChampionship.UpdateOne(context.TODO(), findchampion, update)
		if errUpdateOne != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Participant added successfully",
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server Error",
	})

}
