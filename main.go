package main

import (
	"backend/config"
	"backend/controllers/Post"
	"backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	PORT := config.PORT()

	routes.UseRoutes(app)
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("AAA")
	})
	app.Get("/getPost", Post.GetPost)

	if PORT == "" {
		PORT = "3001"
	}
	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}
