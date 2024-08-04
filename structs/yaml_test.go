package structs_test

import (
	"testing"

	"dac.ac/structs"
	"github.com/stretchr/testify/assert"
)

func TestReadAndUnmarshalYAML(t *testing.T) {
	// Test reading and unmarshalling the articles YAML
	articlesContent := structs.GetArticles()
	assert.NotNil(t, articlesContent, "Expected non-nil output")

	// Test reading and unmarshalling the playlists YAML
	playlistsContent := structs.GetPlaylists()
	assert.NotNil(t, playlistsContent, "Expected non-nil output")

	// Test reading and unmarshalling the links YAML
	linksContent := structs.GetLinks()
	assert.NotNil(t, linksContent, "Expected non-nil output")

	// Test reading and unmarshalling the languages YAML
	languagesContent := structs.GetLanguages()
	assert.NotNil(t, languagesContent, "Expected non-nil output")

	// Test reading and unmarshalling the curriculum YAML
	curriculumContent := structs.GetCurriculum()
	assert.NotNil(t, curriculumContent, "Expected non-nil output")

	// Test reading and unmarshalling the projects YAML
	projectsContent := structs.GetProjects()
	assert.NotNil(t, projectsContent, "Expected non-nil output")
}
