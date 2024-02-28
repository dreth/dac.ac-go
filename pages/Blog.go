package pages

import (
	"strings"

	"dac.ac/components"
	"dac.ac/middleware"
	"dac.ac/structs"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
)

func Blog(c *fiber.Ctx) error {
	// return HTML
	c.Type("html")

	// content
	normalizedPath := strings.TrimSuffix(c.Path(), "/")
	content := components.Content(normalizedPath)

	// Blog page
	html := components.HTML("Daniel Alonso", "Daniel Alonso", middleware.Lang(c), false, content).Render()

	// return nil
	return c.SendString(html)
}

func BlogContent(c *fiber.Ctx) error {
	lang := middleware.Lang(c)
	return c.SendString(BlogContentHTML(lang).Render())
}

func BlogContentHTML(lang string) *elem.Element {
	// Get content from YAML
	// get the article data
	articles := structs.GetArticles()
	text := structs.GetLanguagesLoc(lang)

	// length of arrays of articles
	articleCount := articles.Len()

	// Blog content
	content := elem.Div(
		attrs.Props{
			attrs.Class: "container mx-auto text-left mb-6 pt-3 md:pt-4 xl:pt-5",
		},
		elem.H2(
			attrs.Props{
				attrs.Class: components.RtH("font-black mb-2"),
			},
			elem.Text(text["blog_title"]),
		),
		elem.P(
			nil,
			elem.Text(text["blog_introduction"]),
		),
		elem.Br(nil),
		elem.Div(
			nil,
			ArticleList(articleCount)...,
		),
	)

	// return nil
	return content
}
