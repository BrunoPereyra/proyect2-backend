package main

import (
	"backend/config"
	"backend/database"
	"backend/models"
	"backend/routes"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	PORT := config.PORT()

	routes.UseRoutes(app)
	app.Get("/home", func(c *fiber.Ctx) error {
		db, err := database.GoMongoDB()
		if err != nil {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"message": "StatusServiceUnavailable",
			})
		}
		collection := db.Collection("users")

		// traer primeros campeonatos si esta vacio el el req
		var user models.User
		datafind, _ := primitive.ObjectIDFromHex("6441dee7fa046c203a30b030")
		errfindChampions := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: datafind}}).Decode(&user)
		if errfindChampions != nil {
			fmt.Println("errfindChampions")
		}
		return c.JSON(fiber.Map{
			"data": user,
		})
	})

	if PORT == "" {
		PORT = "3001"
	}
	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}
