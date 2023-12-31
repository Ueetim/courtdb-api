package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ueetim/court-system/database"
	"github.com/ueetim/court-system/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(
		cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins: "http://localhost:4200, https://ecourtpro.vercel.app",
		}),
	)

	// serve static files
	app.Static("/public", "./public/uploads")

	routes.Setup(app)

	log.Fatal(app.Listen(":4000"))
}
