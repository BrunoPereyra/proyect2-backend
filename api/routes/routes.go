package routes

import (
	"backend/api/controllers"
	"backend/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UseRoutes(app *fiber.App) {

	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)

	app.Post("/CreateEvent", middleware.UseExtractor(), controllers.CreateEvent)
}
