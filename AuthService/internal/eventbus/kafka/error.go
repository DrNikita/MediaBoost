package kafka

import (
	"fmt"
)

type Error int

type ErrorWrap struct {
	ErrType Error
	Err     error
}

func (e Error) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e, e.Title(), e.Description())
}

func (e ErrorWrap) Error() string {
	return fmt.Sprintf("[%d] %s: %s: err: %s", e.ErrType, e.ErrType.Title(), e.ErrType.Description(), e.Err.Error())
}

func (e ErrorWrap) Unwrap() error {
	return e.Err
}

const (
	Unknown      Error = -1
	ProtoMarshal Error = 1
	KafkaWrite   Error = 2
)

func (e Error) Title() string {
	switch e {
	case ProtoMarshal:
		return "ProtoMarshal"
	case KafkaWrite:
		return "KafkaWrite"
	case Unknown:
		return "Unknown"
	}
	return "Undefined"
}

func (e Error) Description() string {
	switch e {
	case ProtoMarshal:
		return "failed to marshal protobuf message"
	case KafkaWrite:
		return "failed to write message to Kafka"
	case Unknown:
		return "an unexpected server error occurred"
	}
	return "no description"
}
