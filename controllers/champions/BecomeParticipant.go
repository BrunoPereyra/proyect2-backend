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

type ReqData struct {
	UserID         string `json:"userID"`
	ChampionshipID string `json:"championshipID"`
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

	dataMiddleware := c.Context().UserValue("nameUser")
	dataMiddlewareString, _ := dataMiddleware.(string)
	user, errUserTMiddlExist := helpers.UserTMiddlExist(dataMiddlewareString, db)

	if errUserTMiddlExist != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	idChampionship, errorID := primitive.ObjectIDFromHex(req.ChampionshipID)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id StatusBadRequest",
		})
	}

	var Championship models.Championships
	findchampion := bson.D{
		{Key: "_id", Value: idChampionship},
	}
	CollectionChampionship := db.Collection("championship")
	errFindChampionship := CollectionChampionship.FindOne(context.TODO(), findchampion).Decode(&Championship)
	if errFindChampionship != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Championship not found"})
	}

	UserIDReq, errorID := primitive.ObjectIDFromHex(req.UserID)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id StatusBadRequest",
		})
	}
	if Championship.Creator != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden"})
	}

	update := bson.M{
		"$pull": bson.M{
			"applicants": UserIDReq,
		},
		"$addToSet": bson.M{
			"participants": UserIDReq,
		},
	}

	fmt.Println(UserIDReq)
	_, errUpdateOne := CollectionChampionship.UpdateOne(context.TODO(), findchampion, update)
	if errUpdateOne != nil {
		fmt.Println(errUpdateOne)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{
		"message": "Participant added successfully",
	})
}
