package pages

import "github.com/gofiber/fiber/v2"

func LinkedIn(c *fiber.Ctx) error {
	return c.Redirect("https://linkedin.com/in/dreth")
}

func GitHub(c *fiber.Ctx) error {
	return c.Redirect("https://github.com/dreth")
}

func Email(c *fiber.Ctx) error {
	return c.Redirect("mailto:d@dac.ac")
}
