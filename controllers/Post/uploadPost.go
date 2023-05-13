package Post

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostBody struct {
	Status string `bson:"Status"`
}

func UploadPost(c *fiber.Ctx) error {

	// process image
	fileHeader, _ := c.FormFile("PostImage")
	PostImageChanel := make(chan string)
	errChanel := make(chan error)
	go helpers.Processimage(fileHeader, PostImageChanel, errChanel)

	// database
	db, err := database.NewMongoDB(10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Disconnect(context.Background())
	databaseGoMongodb := db.Pool.Database("goMoongodb")

	// exist user?
	dataMiddleware := c.Context().UserValue("nameUser")
	UserCreator := make(chan models.User)
	errChanelUserTMiddlExist := make(chan error)
	go helpers.UserTMiddlExist(dataMiddleware.(string), databaseGoMongodb, UserCreator, errChanelUserTMiddlExist)

	PostCollection := databaseGoMongodb.Collection("post")

	// validator
	var PostBodyParser PostBody
	errBodyParser := c.BodyParser(&PostBodyParser)
	if errBodyParser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}
	if PostBodyParser.Status == "" || len(PostBodyParser.Status) >= 100 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "es mayor a 100 o ` ` ",
		})
	}

	// PostImageChanel
	for {
		select {
		case PostImage := <-PostImageChanel:
			UserCreatorID := <-UserCreator
			// insert post
			var newPost models.Post
			newPost.UserID = UserCreatorID.ID
			newPost.Status = PostBodyParser.Status
			newPost.PostImage = PostImage
			newPost.TimeStamp = time.Now()
			newPost.Likes = []primitive.ObjectID{}

			postInset, err := PostCollection.InsertOne(context.TODO(), newPost)
			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
					"err":     err,
				})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "StatusOK",
				"data":    postInset,
			})

		case err = <-errChanel:

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "StatusInternalServerError",
			})

		case <-errChanelUserTMiddlExist:
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
				"message": "StatusNotAcceptable",
			})

		}
	}

}
