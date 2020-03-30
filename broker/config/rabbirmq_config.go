package config

import (
	"fmt"
	"github.com/rezwanul-haque/History-Service/utils/helpers"
)

// RabbitMQ ...
type RabbitMQ struct {
	URI string
}

var rabbitMQ = &RabbitMQ{}

// RabbitMQCfg returns configs
func RabbitMQCfg() *RabbitMQ {
	return rabbitMQ
}

// LoadRabbitMQCfg loads configs
func LoadRabbitMQCfg() {
	rabbitMQ.URI = fmt.Sprintf("%s://%s:%s@%s:%d/",
		helpers.GoDotEnvVariable("RABBITMQ_PROTOCOL"),
		helpers.GoDotEnvVariable("RABBITMQ_USERNAME"),
		helpers.GoDotEnvVariable("RABBITMQ_PASSWORD"),
		helpers.GoDotEnvVariable("RABBITMQ_HOST"),
		helpers.GoDotEnvVariable("RABBITMQ_PORT"))
}
