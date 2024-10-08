package main

import (
	"log"

	"dac.ac/server"
)

func main() {
	// run the server
	app := server.ServerWithRender()

	// Splash screen
	println(`
┌┬─┐┌─┐┌─┐   ┌─┐┌─┐
 │ │├─┤│     ├─┤│  
─┴─┘┴ ┴└─┘ o ┴ ┴└─┘`,
	)

	// Start the app
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
