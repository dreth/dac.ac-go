package routes_test

import (
	"testing"

	"dac.ac/server"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// creating fiber app
	app := server.Server()

	// Define the expected routes
	expectedRoutes := []struct {
		method string
		path   string
	}{
		{"GET", "/"},
		{"GET", "/home"},
		{"GET", "/home/content"},
		{"GET", "/cv"},
		{"GET", "/cv/content"},
		{"GET", "/cv-detail"},
		{"GET", "/blog"},
		{"GET", "/blog/content"},
		{"GET", "/cool-links"},
		{"GET", "/cool-links/content"},
		{"GET", "/projects"},
		{"GET", "/projects/content"},
		{"GET", "/playlists"},
		{"GET", "/playlists/content"},
		{"GET", "/profiles/linkedin"},
		{"GET", "/profiles/github"},
		{"GET", "/profiles/email"},
		{"POST", "/toggle-lang"},
		{"GET", "/blog/:article"},
	}

	// Iterate over the expected routes and check if they are registered
	for _, route := range expectedRoutes {
		found := false
		for _, r := range app.GetRoutes() {
			if r.Method == route.method && r.Path == route.path {
				found = true
				break
			}
		}
		assert.True(t, found, "Route not registered: %s %s", route.method, route.path)
	}
}
