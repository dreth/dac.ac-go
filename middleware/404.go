package middleware

import "github.com/gofiber/fiber/v2"

func NotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Redirect to homepage
		return c.Redirect("/", fiber.StatusFound)
	}
}
