package Post

import (
	"backend/database"
	"backend/models"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Idpost struct {
	IDPost string `json:"id_post"`
}

func LikePost(c *fiber.Ctx) error {
	var idPostReq Idpost
	c.BodyParser(&idPostReq)

	idPost, err := primitive.ObjectIDFromHex(idPostReq.IDPost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "StatusBadRequest",
		})
	}
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	collectionPost := databaseGoMongodb.Collection("post")
	findPost := bson.D{
		{Key: "_id", Value: idPost},
	}
	var PostDocument models.Post
	PostCollectionErr := collectionPost.FindOne(context.TODO(), findPost).Decode(&PostDocument)

	if PostCollectionErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "StatusNotFound",
		})
	}

	IDUser := c.Context().UserValue("_id")
	stringIDUser, stringIDUserok := IDUser.(string)
	if !stringIDUserok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "StatusInternalServerError",
		})
	}
	IDUserObject, errinObjectID := primitive.ObjectIDFromHex(stringIDUser)

	if errinObjectID != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "StatusInternalServerError",
		})
	}

	for i, like := range PostDocument.Likes {
		if like == IDUserObject {
			PostDocument.Likes = append(PostDocument.Likes[:i], PostDocument.Likes[i+1:]...)

			update := bson.M{
				"$set": bson.M{
					"Likes": PostDocument.Likes,
				},
			}
			_, errUpdateOne := collectionPost.UpdateOne(context.TODO(), findPost, update)
			if errUpdateOne != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "take out how",
			})
		}
	}

	PostDocument.Likes = append(PostDocument.Likes, IDUserObject)
	update := bson.M{
		"$set": bson.M{
			"Likes": PostDocument.Likes,
		},
	}
	_, errUpdateOne := collectionPost.UpdateOne(context.TODO(), findPost, update)
	if errUpdateOne != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Like",
	})
}
