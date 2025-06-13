package kafka

import (
	"auth-service/config"
	"context"

	pb "auth-service/internal/eventbus/gen/eventbus"

	"github.com/segmentio/kafka-go"

	"google.golang.org/protobuf/proto"
)

type Eventbus struct {
	writer *kafka.Writer
	config config.EventbusConfig
}

func NewEventbus(config config.EventbusConfig) Eventbus {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: config.BrokerAddresses,
		Topic:   config.TopicRegistration,
	})

	return Eventbus{
		writer: writer,
		config: config,
	}
}

func (e *Eventbus) NotifyRegistration(ctx context.Context, userId int64, email string) error {
	messageBytes, err := proto.Marshal(&pb.UserRegistrated{
		UserId: userId,
		Email:  email,
	})
	if err != nil {
		return ErrorWrap{
			ErrType: Unknown,
			Err:     err,
		}
	}

	err = e.writer.WriteMessages(ctx, kafka.Message{
		Value: messageBytes,
	})
	if err != nil {
		return ErrorWrap{
			ErrType: Unknown,
			Err:     err,
		}
	}

	return nil
}
