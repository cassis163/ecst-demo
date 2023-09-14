package main

import (
	"ecst-demo/cassis163/article-service/backend"
	"ecst-demo/cassis163/article-service/frontend"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./frontend/templates", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	backend.InitBackend(app)
	frontend.InitWeb(app)

	log.Fatal(app.Listen(":3000"))
}
