package server

import (
	"dac.sg/middleware"
	"dac.sg/routes"
	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	// Create Fiber app
	app := fiber.New()

	// Other middlewares
	app.Use(middleware.Language())

	// Serve static files
	app.Static("/", "./static", fiber.Static{
		MaxAge:   31536000,
		Compress: true,
	})

	// Register routes
	routes.Routes(app)

	// Handle not found errors
	app.Use(middleware.NotFound())

	// Return the configured app
	return app
}

func ServerWithRender() *fiber.App {
	// Render components (assuming RenderComponents is defined elsewhere)
	RenderComponents()

	// Create and return the server
	return Server()
}
