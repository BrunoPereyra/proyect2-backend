package Post

import (
	"backend/config"
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/validator"
	"context"
	"mime/multipart"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadPost(c *fiber.Ctx) error {
	// database
	Database, errDB := database.GoMongoDB()
	if errDB != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})

	}

	// process image

	fileHeader, _ := c.FormFile("PostImage")
	PostImageChanel := make(chan string)
	errChanel := make(chan error)
	go Processimage(fileHeader, PostImageChanel, errChanel)

	// validator
	var PostBodyParser validator.UploadPostValidate
	error := c.BodyParser(&PostBodyParser)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
		})
	}

	UploadPostValidateErr := PostBodyParser.UploadPostValidate()
	if UploadPostValidateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"messages": "Bad Request",
			"data":     UploadPostValidateErr,
		})
	}
	// usuario existe?
	dataMiddleware := c.Context().UserValue("nameUser")
	dataMiddlewareString, _ := dataMiddleware.(string)

	UserCreator, err := helpers.UserTMiddlExist(dataMiddlewareString, Database)
	if err != nil {
		return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	// crear

	// PostImageChanel
	for {
		select {
		case PostImage := <-PostImageChanel:
			// insert post
			var newPost models.Post
			newPost.UserID = UserCreator.ID
			newPost.Status = PostBodyParser.Status
			newPost.PostImage = PostImage
			newPost.TimeStamp = time.Now()
			newPost.Likes = []primitive.ObjectID{}

			PostCollection := Database.Collection("post")
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
		}

	}
}
func Processimage(fileHeader *multipart.FileHeader, PostImageChanel chan string, errChanel chan error) {
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()
		ctx := context.Background()
		cldService, errcloudinary:= cloudinary.NewFromURL(config.CLOUDINARY_URL())
		if errcloudinary != nil {
			rrChanel <- errcloudinary
		}
	resp, errcldService := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		if errcldService != nil || strings.HasPrefix(resp.SecureURL, "https://") {
			rrChanel <- errcldService
	}

		PostImaeChanel <- resp.SecureURL
	} else {
		ostImageChanel <- ""
	
}
