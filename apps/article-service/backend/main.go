package backend

import (
	"github.com/gofiber/fiber"
)

func InitBackend(app *fiber.App) {
	InitOrm()
	InitApi(app)
}
