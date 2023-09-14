package backend

import "github.com/gofiber/fiber"

func InitApi(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) {
		c.SendString("Hello, World!")
	})
}
