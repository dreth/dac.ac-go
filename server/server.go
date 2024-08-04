package server

import (
	"dac.ac/middleware"
	"dac.ac/routes"
	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	// creating fiber app
	app := fiber.New()

	// Middlewares
	app.Use(middleware.Language())

	// Serve static files
	app.Static("/", "./static")

	// routes
	routes.Routes(app)

	// Errors
	app.Use(middleware.NotFound())

	// return the app
	return app
}

func ServerWithRender() *fiber.App {
	// Render components
	RenderComponents()

	// Create the app
	return Server()
}
