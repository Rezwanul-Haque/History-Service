package consumer

import (
	"github.com/rezwanul-haque/History-Service/clients/config"
)

func Consumer() {
	connectionString := config.GetConnectionString()
	config.MessageQueueService.ConnectToBroker(connectionString)
	queueName := config.GetQueueName()

	config.MessageQueueService.SubscribeToQueue(queueName, "")
}
