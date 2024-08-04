package main_test

import (
	"net/http"
	"testing"
	"time"

	"dac.ac/server"
)

func TestServer(t *testing.T) {
	// Create a new fiber app
	app := server.Server()

	// Run the app in a separate goroutine
	go func() {
		if err := app.Listen(":7741"); err != nil {
			t.Error("Failed to start server:", err)
		}
	}()

	// Wait for the app to start listening
	time.Sleep(1 * time.Second)

	// Make a GET request to the root URL
	resp, err := http.Get("http://localhost:7741")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Stop the app
	app.Shutdown()
}
