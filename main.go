package main

import (
	"backend/config"
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

	if PORT == "" {
		PORT = "3001"
	}
	log.Fatal(app.Listen(PORT))
}
