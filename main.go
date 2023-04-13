package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	PORT := config.PORT()
	routes.UseRoutes(app)
	app.Listen(PORT)
}
