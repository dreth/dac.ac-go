package structs

import (
	"fmt"
	"os"
	"path/filepath"

	"dac.ac/utils"
	orderedmap "github.com/wk8/go-ordered-map/v2"

	"gopkg.in/yaml.v3"
)

func ReadAndUnmarshalYAML(fileName string, out interface{}) {
	// Read the yaml file into a byte array
	data, err := os.ReadFile(filepath.Join("data", fileName))
	utils.HandleErr(err, false, fmt.Sprintf("error reading %s", fileName))

	// Unmarshal the yaml into the provided struct
	err = yaml.Unmarshal(data, out)
	utils.HandleErr(err, false, fmt.Sprintf("error unmarshalling %s", fileName))
}

func GetArticles() orderedmap.OrderedMap[string, Article] {
	var content orderedmap.OrderedMap[string, Article]
	ReadAndUnmarshalYAML("articles.yml", &content)
	return content
}

func GetPlaylists() orderedmap.OrderedMap[string, string] {
	var content orderedmap.OrderedMap[string, string]
	ReadAndUnmarshalYAML("playlists.yml", &content)
	return content
}

func GetLinks() Links {
	var content Links
	ReadAndUnmarshalYAML("links.yml", &content)
	return content
}

func GetLanguages() Languages {
	var content Languages
	ReadAndUnmarshalYAML("languages.yml", &content)
	return content
}

func GetCurriculum() CV {
	var content CV
	ReadAndUnmarshalYAML("curriculum.yml", &content)
	return content
}

func GetLanguagesLoc(lang string) map[string]string {
	langs := GetLanguages()
	if lang == "es" {
		return langs.Es
	}
	return langs.En
}

func GetCurriculumLoc(lang string) Language {
	curriculum := GetCurriculum()
	if lang == "es" {
		return curriculum.Es
	}
	return curriculum.En
}
