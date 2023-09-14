package backend

import "github.com/gofiber/fiber/v2"

func InitApi(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
