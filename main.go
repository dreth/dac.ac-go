package main

import (
	"log"

	"dac.sg/server"
)

func main() {
	// run the server
	app := server.ServerWithRender()

	// Start the app
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
