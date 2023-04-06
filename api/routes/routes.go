package routes

import (
	"backend/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UseRoutes(app *fiber.App) {
	app.Post("/signup", controllers.Signup)
}
