package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rezwanul-haque/History-Service/clients/consumer"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	consumer.Consumer()

	router.Run(":8090")
}
