package consumer

import (
	"errors"
	"github.com/rezwanul-haque/History-Service/utils/consts"
)

var (
	// ErrInvalidQueueType ...
	ErrInvalidQueueType = errors.New("invalid consumer type")
)

// Connect ...
func Connect(qt consts.QueueType) error {
	if qt == consts.RabbitMQ {
		return ConnectRabbitMQ()
	}
	return ErrInvalidQueueType
}

// GetConnection ...
func GetConnection(qt consts.QueueType) interface{} {
	if qt == consts.RabbitMQ {
		return GetRabbitMQ()
	}
	return ErrInvalidQueueType
}
