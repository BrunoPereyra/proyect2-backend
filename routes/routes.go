package routes

import (
	"backend/controllers"
	"backend/controllers/Post"
	"backend/controllers/champions"
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func UseRoutes(app *fiber.App) {
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)

	app.Get("/Currentuser", middleware.UseExtractor(), controllers.Currentuser)
	app.Post("/Searchuser", middleware.UseExtractor(), controllers.SearchUser)
	//post
	app.Post("/UploadPost", middleware.UseExtractor(), Post.UploadPost)
	app.Get("/getPost", Post.GetPost)

	// ----- eventos ----
	// user
	app.Post("/CreateEvent", middleware.UseExtractor(), champions.CreateChampionship)

	app.Post("/SearchChampions", middleware.UseExtractor(), champions.SearchChampions)
	// auto
	app.Post("/getEvent", champions.GetChampionshipSID)
	app.Get("/getEvent", champions.GetChampionships)
}
