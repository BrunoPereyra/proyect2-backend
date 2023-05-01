package champions

import (
	"backend/database"
	"backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Idchampionship struct {
	IDchampionship string `json:"IDchampionship"`
}

func AskForChampionship(c *fiber.Ctx) error {
	// connect database
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	var req Idchampionship
	if errBodyparser := c.BodyParser(&req); errBodyparser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}

	collectionchampionship := databaseGoMongodb.Collection("championship")
	id, errorID := primitive.ObjectIDFromHex(req.IDchampionship)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}

	find := bson.D{
		{Key: "_id", Value: id},
	}
	var collectionChampionship models.Championships
	errfindChampion := collectionchampionship.FindOne(context.TODO(), find).Decode(&collectionChampionship)
	if errfindChampion != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	dataMiddleware := c.Context().UserValue("_id")
	dataMiddlewareObjectID, errdataMiddlewareObjectID := primitive.ObjectIDFromHex(dataMiddleware.(string))
	if errdataMiddlewareObjectID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest creator error",
		})
	}
	if dataMiddlewareObjectID == collectionChampionship.Creator {
		var acceptedUsers []models.User

		// Obtener los detalles de los usuarios con los IDs en acceptedApplicants
		usersCollection := databaseGoMongodb.Collection("users")
		usersCursor, err := usersCollection.Find(context.TODO(), bson.M{
			"_id": bson.M{
				"$in": collectionChampionship.AcceptedApplicants,
			},
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error retrieving users",
			})
		}
		for usersCursor.Next(context.TODO()) {
			var user models.User
			if err := usersCursor.Decode(&user); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Error decoding user",
				})
			}
			// Agregar el usuario a la lista de usuarios aceptados
			acceptedUsers = append(acceptedUsers, user)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":            "ok",
			"data":               collectionChampionship,
			"acceptedApplicants": acceptedUsers,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
			"data":    collectionChampionship,
		})
	}

	// Actualizar el campo acceptedApplicants de la instancia collectionChampionship

}
