package utils

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetLang(c *fiber.Ctx) string {
	return c.Locals("lang").(string)
}

// Get lang cookie
func GetLangCookie(c *fiber.Ctx) string {
	return c.Cookies("language")
}

// Set lang cookie
func SetLangCookie(c *fiber.Ctx, lang string) {
	c.Cookie(&fiber.Cookie{
		Name:    "language",
		Value:   lang,
		Expires: time.Now().Add(720 * time.Hour), // Expires in 30 days
	})
}

// Get CV expanded/collapsed status cookie
func GetCVDetailCookie(c *fiber.Ctx) string {
	return c.Cookies("cv-detail")
}

// Set CV expanded/collapsed status cookie
func SetCVDetailCookie(c *fiber.Ctx, state string) {
	c.Cookie(&fiber.Cookie{
		Name:    "cv-detail",
		Value:   state,
		Expires: time.Now().Add(168 * time.Hour), // Expires in 7 days
	})
}

// helper func to replace spaces with dashes
func RSWD(s string) string {
	return strings.ReplaceAll(s, " ", "-")
}

// handle error func
func HandleErr(err error, fatal bool, msg string) error {
	// log error and exit if fatal is true
	if err != nil {
		if !fatal {
			log.Println(err)
			log.Println(msg)
		} else {
			log.Fatal(err)
			log.Fatal(msg)
		}
		return err
	}

	// return nil if no error
	return nil
}

// handle error func
func HandleErrBool(err bool, fatal bool, msg string) bool {
	// log error and exit if fatal is true
	if !err {
		if !fatal {
			log.Println(err)
			log.Println(msg)
		} else {
			log.Fatal(err)
			log.Fatal(msg)
		}
		return err
	}

	// return false if no error
	return false
}
