package middleware

import (
	"dac.sg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

func Language() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Use the Lang function to determine the language.
		lang := Lang(c)
		// Set the language in the locals storage for the next handlers.
		c.Locals("lang", lang)
		return c.Next()
	}
}

func Lang(c *fiber.Ctx) string {
	// Define the supported languages.
	supportedLangs := []language.Tag{
		language.English, // en
		language.Spanish, // es
	}

	// Use a matcher to find the best matching language.
	matcher := language.NewMatcher(supportedLangs)

	// Check 'language' cookie first
	if lang := utils.GetLangCookie(c); lang != "" {
		// Validate if the language is supported.
		if _, _, confidence := matcher.Match(language.Make(lang)); confidence != language.No {
			return lang
		}
	}

	// Parse the Accept-Language header.
	accept := c.Get("Accept-Language")
	tags, _, err := language.ParseAcceptLanguage(accept)
	if err == nil {
		// If the header is parsed successfully, find the most preferred language.
		tag, _, _ := matcher.Match(tags...)
		base, _ := tag.Base()
		return base.String()
	}

	// Default to 'en' if no valid language is found.
	return "en"
}

func NextLang(lang string) string {
	return map[string]string{
		"en": "es",
		"es": "en",
	}[lang]
}
