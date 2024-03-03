package components

import (
	"dac.ac/middleware"
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/gofiber/fiber/v2"
)

func HTMLProps() attrs.Props {
	return attrs.Props{
		attrs.Class: "bg-background text-text fonts",
	}
}

func HTML(Title string, Description string, lang string, Article bool, content *elem.Element) *elem.Element {
	return elem.Html(
		HTMLProps(),
		Head(Title, Description, Article),
		Body(content, lang),
	)
}

func PageHTML(c *fiber.Ctx, content *elem.Element) error {
	// return HTML
	c.Type("html")

	// Homepage
	html := HTML("Daniel Alonso", "Daniel Alonso", middleware.Lang(c), false, Content(content)).Render()

	// return nil
	return c.SendString(html)
}
