package payment

import (
	"backend/database"
	"backend/models"
	"context"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

func TicketPaymentConfimation(c *fiber.Ctx) error {
	// crear en el front para la solicitud
	db, err := database.GoMongoDB()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	collection := db.Collection("championship")
	find := bson.D{
		{Key: "_id", Value: "6449c51d217cb4159cfa1eb2"},
	}
	var collectionChampionship models.Championships
	errfindChampion := collection.FindOne(context.TODO(), find).Decode(&collectionChampionship)

	if errfindChampion != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	// ver si el usuario fue aceptado
	// for _, v := range collectionChampionship.AcceptedApplicants {
	// if v == idUser{

	// }
	// }

	// update := bson.M{
	// 	"$pull": bson.M{
	// 		"applicants": UserIDReq,
	// 	},
	// 	"$set": bson.M{
	// 		"participantesquepagaronlaentrada": Championship.AcceptedApplicants,
	// 	},
	// }

	//errUpdateOne := collection.UpdateOne(context.TODO(), find, update)
	// if errUpdateOne != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Internal Server Error",
	// 	})
	// }
	return err
}
