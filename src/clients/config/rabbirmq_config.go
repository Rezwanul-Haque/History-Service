package config

import (
	"encoding/json"
	"fmt"
	"github.com/rezwanul-haque/History-Service/src/domain/queue"
	"github.com/rezwanul-haque/History-Service/src/logger"
	"github.com/rezwanul-haque/History-Service/src/services"
	"github.com/rezwanul-haque/History-Service/src/utils/helpers"
	"github.com/streadway/amqp"
	"os"
)

var (
	MessageQueueService IMessagingClient = &MessagingClient{}
)

// Defines our interface for connecting and consuming messages.
type IMessagingClient interface {
	ConnectToBroker(string)
	SubscribeToQueue(queueName string, consumerName string) error
	Close()
}

// Real implementation, encapsulates a pointer to an amqp.Connection
type MessagingClient struct {
	conn *amqp.Connection
}

func (m *MessagingClient) ConnectToBroker(connectionString string) {
	if connectionString == "" {
		panic("Cannot initialize connection to broker, connectionString not set. Have you initialized?")
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", connectionString))
	if err != nil {
		logger.Error("Failed to connect to AMQP compatible broker at: ", err)
	}
}

// get connection String
func GetConnectionString() string {
	connectionString := fmt.Sprintf("%s://%s:%s@%s/",
		helpers.GoDotEnvVariable("RABBITMQ_PROTOCOL"),
		helpers.GoDotEnvVariable("RABBITMQ_USERNAME"),
		helpers.GoDotEnvVariable("RABBITMQ_PASSWORD"),
		helpers.GoDotEnvVariable("RABBITMQ_HOST"))
	//helpers.GoDotEnvVariable("RABBITMQ_PORT")) // Need this for run without docker

	return connectionString
}

func GetQueueName() string {
	return helpers.GoDotEnvVariable("RABBITMQ_QUEUE_NAME")
}

func (m *MessagingClient) SubscribeToQueue(queueName string, consumerName string) error {
	ch, err := m.conn.Channel()
	failOnError(err, "Failed to open a channel")

	logger.Info(fmt.Sprintf("Declaring Queue (%s)", queueName))
	queue, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	failOnError(err, "Failed to register an Queue")

	msgs, err := ch.Consume(
		queue.Name,   // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	go consumeLoop(msgs, acknowledgeMessages)
	return nil
}

func (m *MessagingClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func consumeLoop(deliveries <-chan amqp.Delivery, handlerFunc func(d amqp.Delivery)) {
	for d := range deliveries {
		// Invoke the handlerFunc func we passed as parameter.
		handlerFunc(d)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func acknowledgeMessages(d amqp.Delivery) {
	logger.Info(fmt.Sprintf("Consumer ready, PID: %d", os.Getpid()))
	logger.Info(fmt.Sprintf("Received a message: %s", string(d.Body)))
	locationData := &queue.RabbitLocationData{}

	err := json.Unmarshal(d.Body, locationData)

	if err != nil {
		logger.Error("Error decoding JSON: %s", err)
		return
	}

	logger.Info("Result: %s" + locationData.ToString())

	_, getErr := services.HistoryService.CreateHistory(*locationData)
	if getErr != nil {
		logger.Error("save to database failed: %s", err)
		return
	}

	if err := d.Ack(false); err != nil {
		logger.Error("Error acknowledging message: %s", err)
		return
	} else {
		logger.Info("Acknowledged message")
	}
}
