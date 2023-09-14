package backend

import (
	"github.com/gofiber/fiber/v2"
)

func InitBackend(app *fiber.App) {
	InitOrm()
	InitApi(app)
}
