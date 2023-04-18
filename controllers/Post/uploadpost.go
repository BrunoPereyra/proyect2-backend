package Post

import (
	"backend/config"
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/validator"
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

func UploadPost(c *fiber.Ctx) error {
	// database
	Database, errDB := database.GoMongoDB()
	if errDB != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"message": "StatusServiceUnavailable",
		})

	}

	fileHeader, _ := c.FormFile("PostImage")

	var PostImage string
	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()
		cldService, _ := cloudinary.NewFromURL(config.CLOUDINARY_URL())
		resp, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
		PostImage = resp.SecureURL
	} else {
		PostImage = ""
	}

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

	var newPost models.Post
	newPost.UserID = UserCreator.ID
	newPost.Status = PostBodyParser.Status
	newPost.PostImage = PostImage
	newPost.TimeStamp = time.Now()

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
}
