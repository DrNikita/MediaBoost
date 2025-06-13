package kafka

import (
	"auth-service/config"
	"context"
	"encoding/json"

	pb "auth-service/internal/eventbus/gen/eventbus"

	"github.com/segmentio/kafka-go"
)

var topicNames map[string]string

func init() {
	topicNames = map[string]string{
		"registration": "user.registration",
	}
}

type Eventbus struct {
	config config.EventbusConfig
}

func NewEventbus(config config.EventbusConfig) Eventbus {
	return Eventbus{
		config: config,
	}
}

func (e *Eventbus) NewwWriter() *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: e.config.BrokerAddresses,
		Topic:   topicNames["registration"],
	})

	return writer
}

func (e *Eventbus) NotifyRegisteration(ctx context.Context, userId string) error {
	writer := e.NewwWriter()
	defer writer.Close()

	message := pb.UserRegistrated{
		UserId: 0,
		Email:  "",
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return Unknown
	}

	writer.WriteMessages(ctx, kafka.Message{
		Value: messageBytes,
	})

	return nil
}
