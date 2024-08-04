package pages

import (
	"dac.ac/components"
	"dac.ac/middleware"
	"dac.ac/structs"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
)

func Blog(c *fiber.Ctx) error {
	// return the full page
	return components.PageHTML(c, BlogContentHTML(middleware.Lang(c)))
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
				elem.Text(text["blog_title"]),
			),
			elem.P(
				nil,
				elem.Text(text["blog_introduction"]),
			),
		),
		elem.Div(
			attrs.Props{
				attrs.Class: "py-2 space-y-2",
			},
			ArticleList(articleCount)...,
		),
	)

	// return nil
	return content
}
