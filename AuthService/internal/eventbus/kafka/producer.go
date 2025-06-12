package kafka

import (
	"auth-service/config"
	"context"

	"github.com/segmentio/kafka-go"
)

type Eventbus struct {
	config config.EventbusConfig
}

func NewEventbus(config config.EventbusConfig) Eventbus {
	return Eventbus{
		config: config,
	}
}

func NewwWriter() *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
	})

	return writer
}

func (e *Eventbus) NotifyRegisteration(ctx context.Context) error {
	return nil
}
