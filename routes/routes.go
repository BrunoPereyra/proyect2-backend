package routes

import (
	"backend/controllers"
	"backend/controllers/champions"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func UseRoutes(app *fiber.App) {
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)
	app.Get("/Currentuser", middleware.UseExtractor(), controllers.Currentuser)
	//post
	app.Post("/UploadPost", middleware.UseExtractor(), controllers.UploadPost)
	// ----- eventos ----
	// user
	app.Post("/CreateEvent", middleware.UseExtractor(), champions.CreateChampionship)
	app.Post("/SearchChampions", middleware.UseExtractor(), champions.SearchChampions)
	// auto
	app.Post("/getEvent", champions.GetChampionshipSID)
	app.Get("/getEvent", champions.GetChampionships)
}
