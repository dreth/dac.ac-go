package routes

import (
	"dac.ac/middleware"
	"dac.ac/pages"
	"dac.ac/structs"
	"dac.ac/utils"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Homepage
	app.Get("/", pages.Homepage)
	app.Get("/home", pages.Homepage)
	app.Get("/home/content", pages.HomepageContent)

	// CV/About
	app.Get("/cv", pages.CV)
	app.Get("/cv/content", pages.CVContent)
	app.Get("/cv-detail", pages.CVDetail)

	// Blog
	app.Get("/blog", pages.Blog)
	app.Get("/blog/content", pages.BlogContent)

	// Cool links
	app.Get("/cool-links", pages.CoolLinks)
	app.Get("/cool-links/content", pages.CoolLinksContent)

	// Playlists
	app.Get("/playlists", pages.Playlists)
	app.Get("/playlists/content", pages.PlaylistsContent)

	// Social profiles
	app.Get("/profiles/linkedin", pages.LinkedIn)
	app.Get("/profiles/github", pages.GitHub)
	app.Get("/profiles/x", pages.X)
	app.Get("/profiles/email", pages.Email)

	// toggle language
	app.Post("/toggle-lang", ToggleLang)

	// Articles
	app.Get("/blog/:article", pages.Article)
}

func ToggleLang(c *fiber.Ctx) error {
	// Attempt to parse `nextLang` from the request body
	var body struct {
		NextLang string `json:"lang"`
	}
	var lang string = "en"
	// If parsing fails, fallback to toggling language based on the current setting
	if err := c.BodyParser(&body); err != nil || body.NextLang == "" {
		lang = middleware.Lang(c) // Assuming Lang(c) is a function to get the current language
		body.NextLang = middleware.NextLang(lang)
	}

	// Get the new language path and button text
	text := structs.GetLanguagesLoc(lang)
	langText := map[string]string{
		"es": text["lang_button_en"],
		"en": text["lang_button_es"],
	}[body.NextLang]

	// Set the new language in the cookie
	utils.SetLangCookie(c, body.NextLang)

	// Set the new language in c.Locals for immediate access
	return c.SendString(langText)
}
