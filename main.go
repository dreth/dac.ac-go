package main

import (
	"dac.ac/middleware"
	"dac.ac/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// creating fiber app
	app := fiber.New()

	// Middlewares
	app.Use(middleware.Language())

	// Static files
	app.Static("/", "./static")

	// render components
	RenderComponents()

	// routes
	routes.Routes(app)

	// start app
	app.Listen(":8080")
}
