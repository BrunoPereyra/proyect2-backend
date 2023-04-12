package main

import (
	"backend/api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error")
	}
	app := fiber.New()
	app.Use(cors.New())
	PORT := os.Getenv("PORT")
	routes.UseRoutes(app)
	app.Listen(PORT)
}
