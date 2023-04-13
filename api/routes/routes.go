package routes

import (
	"backend/api/controllers"
	"backend/api/controllers/champions"
	"backend/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UseRoutes(app *fiber.App) {

	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)

	// ----- eventos ----
	// user
	app.Post("/CreateEvent", middleware.UseExtractor(), champions.CreateChampionship)
	app.Post("/SearchChampions", middleware.UseExtractor(), champions.SearchChampions)
	// auto
	app.Post("/getEvent", champions.GetChampionshipSID)
	app.Get("/getEvent", champions.GetChampionships)
}
