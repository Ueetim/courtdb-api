package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ueetim/court-system/controllers"
)

func Setup(app *fiber.App) {
	// user auth routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.GetLoggedInUser)
	app.Post("/api/logout", controllers.Logout)

	// record routes
	app.Post("/api/record", controllers.CreateRecord)
	app.Get("/api/records", controllers.GetRecordsByUser)
}