package backend

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
)

func InitBackend(app *fiber.App, kafkaProducer *kafka.Producer) {
	InitOrm()
	InitApi(app, kafkaProducer)
}
