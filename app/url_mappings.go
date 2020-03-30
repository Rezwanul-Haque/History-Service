package app

import (
	"github.com/rezwanul-haque/History-Service/controllers/history"
	"github.com/rezwanul-haque/History-Service/controllers/ping"
	"github.com/rezwanul-haque/History-Service/utils/consts"
)

const (
	APIBase = "api/" + consts.APIVersion
)

func mapUrls() {
	router.GET(APIBase+"/ping", ping.Ping)

	router.GET(APIBase+"/history", history.Get)
}
