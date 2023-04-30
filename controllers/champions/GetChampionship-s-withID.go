package champions

import (
	"backend/database"
	"backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dataGetChampionshipStruct struct {
	ID string `json:"idLastChampionship"`
}

func GetChampionshipSID(c *fiber.Ctx) error {

	var dataGetChampionship dataGetChampionshipStruct
	if err := c.BodyParser(&dataGetChampionship); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	// connect database
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	collection := databaseGoMongodb.Collection("championship")
	// traer primeros campeonatos si esta vacio el el req
	id, errorID := primitive.ObjectIDFromHex(dataGetChampionship.ID)
	if errorID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id unrecognized",
		})
	}

	// buscar id
	var IDCreatedAt models.Championships
	errIDCreatedAt := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&IDCreatedAt)
	if errIDCreatedAt != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "championship unrecognized",
		})
	}

	// buscar 10 adelante de id
	findNextTheId := bson.M{
		"createdat": bson.M{"$gt": IDCreatedAt.CreatedAt},
	}
	options := options.Find().
		SetSort(bson.M{"createdat": 1}).
		SetLimit(3)

	cursor, errfindNextTheId := collection.Find(context.TODO(), findNextTheId, options)
	if errfindNextTheId != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "NotFound more championship",
		})
	}

	// Iterar el cursor para obtener los documentos paginados
	var championships []models.Championships
	for cursor.Next(context.TODO()) {
		var championship models.Championships
		err := cursor.Decode(&championship)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error server",
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
