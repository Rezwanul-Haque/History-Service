package consumer

import (
	"fmt"
	"github.com/rezwanul-haque/History-Service/logger"
	"github.com/rezwanul-haque/History-Service/utils/helpers"
	"github.com/streadway/amqp"
	"log"
)

var (
	protocol = helpers.GoDotEnvVariable("RABBITMQ_PROTOCOL")
	username = helpers.GoDotEnvVariable("RABBITMQ_USERNAME")
	password = helpers.GoDotEnvVariable("RABBITMQ_PASSWORD")
	host     = helpers.GoDotEnvVariable("RABBITMQ_HOST")
	port     = helpers.GoDotEnvVariable("RABBITMQ_PORT")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func initCounsumer() {
	brokerSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/",
		protocol, username, password, host, port,
	)

	conn, err := amqp.Dial(brokerSourceName)
	failOnError(err, "Failed to connect to RabbitMQ consumer")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"location-data-queue", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	failOnError(err, "Failed to declare a consumer")

	msgs, err := ch.Consume(
		q.Name, // consumer
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	logger.Info("rabbitmq successfully configured")
}
