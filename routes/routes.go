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
	app.Post("/Searchuser", controllers.SearchUser)
	//post
	app.Post("/UploadPost", middleware.UseExtractor(), Post.UploadPost)
	app.Get("/getPost", Post.GetPost)
	app.Post("/likePost", middleware.UseExtractor(), Post.LikePost)
	// ----- eventos ----
	app.Post("/CreateEvent", middleware.UseExtractor(), champions.CreateChampionship)
	app.Post("/SearchChampions", middleware.UseExtractor(), champions.SearchChampions)

	app.Post("/ApplyChampionship", middleware.UseExtractor(), champions.ApplyChampionship)
	app.Post("/BecomeParticipant", middleware.UseExtractor(), champions.BecomeParticipant)

	// auto
	app.Post("/getEvent", champions.GetChampionshipSID)
	app.Get("/getEvent", champions.GetChampionships)
}
