package main

import (
	"ecst-demo/cassis163/article-service/backend"
	"ecst-demo/cassis163/article-service/frontend"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./frontend/templates", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	kafkaProducer := initKafkaProducer()
	backend.InitBackend(app, kafkaProducer)
	frontend.InitWeb(app)

	app.Hooks().OnShutdown(func() error {
		kafkaProducer.Close()
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}

func initKafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})
	if err != nil {
		panic(err)
	}

	return p
}
