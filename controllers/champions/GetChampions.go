package champions

import (
	"backend/database"
	"backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetChampionships(c *fiber.Ctx) error {

	// connect database
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	collection := databaseGoMongodb.Collection("championship")

	// traer primeros campeonatos si esta vacio el el req
	options := options.Find().
		SetSort(bson.M{"createdat": 1}).
		SetLimit(3)

	cursor, errfindChampions := collection.Find(context.TODO(), bson.M{}, options)

	if errfindChampions != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})
	}
	// Iterar el cursor para obtener los documentos paginados
	var championships []models.Championships
	for cursor.Next(context.TODO()) {
		var championship models.Championships
		err := cursor.Decode(&championship)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "StatusServiceUnavailable",
			})
		}
		championships = append(championships, championship)
	}

	// Devolver la respuesta de la funcion
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":    championships,
	})
}
