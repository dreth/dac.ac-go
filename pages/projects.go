package pages

import (
	"fmt"

	"dac.sg/components"
	"dac.sg/middleware"
	"dac.sg/structs"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
)

func Projects(c *fiber.Ctx) error {
	// return the full page
	return components.PageHTML(c, ProjectsContentHTML(middleware.Lang(c)))
}

func ProjectsContent(c *fiber.Ctx) error {
	lang := middleware.Lang(c)
	return c.SendString(ProjectsContentHTML(lang).Render())
}

func ProjectList(lang string, amount int) []elem.Node {
	// Get content from YAML
	projects := structs.GetProjectsLoc(lang)

	// Projects block
	var projectBlock []elem.Node

	// values of the articlesList
	projectKeys := make([]string, 0, amount)
	projectValues := make([]structs.ProjectDetail, 0, amount)

	// iterate over project
	cnt := 0
	for pair := projects.Oldest(); pair != nil && cnt < amount; pair = pair.Next() {
		projectKeys = append(projectKeys, pair.Key)
		projectValues = append(projectValues, pair.Value)

		cnt++
	}

	// Projects content
	for i := 0; i < len(projectKeys); i++ {
		// get the project info
		project := projectValues[i]

		// If the URL isn't an empty string, add the emoji to it
		projectUrlBlock := elem.Span(nil)
		if project.Url != "" {
			projectUrlBlock = elem.A(
				attrs.Props{
					attrs.Href:  fmt.Sprintf("https://%s", project.Url),
					attrs.Class: "text-links hover:text-links-hover",
				},
				elem.Text(fmt.Sprintf("🔗 %s &nbsp;&nbsp;&nbsp;", project.Url)),
			)
		}
		// If the repo isn't an empty string, add the emoji to it
		projectRepoBlock := elem.Span(nil)
		if project.Repo != "" {
			projectRepoBlock = elem.A(
				attrs.Props{
					attrs.Href:  fmt.Sprintf("https://%s", project.Repo),
					attrs.Class: "text-links hover:text-links-hover",
				},
				elem.Text("📦 repo"),
			)

		}

		// Project block
		projectBlock = append(projectBlock, elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-4 p-4",
			},
			elem.H3(
				attrs.Props{
					attrs.ID:    "project-heading",
					attrs.Class: components.RtH("font-black"),
				},
				elem.Text(fmt.Sprintf("%s %s", project.Emoji, project.Name)),
			),
			elem.P(
				attrs.Props{
					attrs.Class: "text-left py-2",
				},
				elem.Text(project.Description),
			),
			projectUrlBlock,
			projectRepoBlock,
		))
	}

	return projectBlock
}

func ProjectsContentHTML(lang string) *elem.Element {
	// Get content from YAML
	projects := structs.GetProjectsLoc(lang)
	text := structs.GetLanguagesLoc(lang)

	// Get length of projects
	projectsCount := projects.Len()

	// Projects content
	content := elem.Div(
		attrs.Props{
			attrs.Class: "container mx-auto text-left pb-4 pt-4 sm:pt-4",
		},
		elem.Div(
			attrs.Props{
				attrs.Class: "rounded-lg bg-almost-black-2 p-4 sm:pb-6 pt-4",
			},
			elem.H2(
				attrs.Props{
					attrs.Class: components.RtH("font-black mb-2"),
				},
				elem.Text(text["projects_title"]),
			),
			elem.P(
				nil,
				elem.Text(text["projects_introduction"]),
			),
		),
		elem.Div(
			attrs.Props{
				attrs.Class: "py-2 space-y-2",
			},
			ProjectList(lang, projectsCount)...,
		),
	)

	// return nil
	return content
}
