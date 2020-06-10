package consumer

import (
	"github.com/rezwanul-haque/History-Service/src/clients/config"
	"github.com/rezwanul-haque/History-Service/src/logger"
)

func Consumer() {
	connectionString := config.GetConnectionString()
	logger.Info(connectionString)
	config.MessageQueueService.ConnectToBroker(connectionString)
	queueName := config.GetQueueName()

	config.MessageQueueService.SubscribeToQueue(queueName, "")
}
