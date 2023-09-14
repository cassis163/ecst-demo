package frontend

import (
	"github.com/gofiber/fiber/v2"
)

func InitWeb(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
}
