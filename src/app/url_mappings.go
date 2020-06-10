package app

import (
	"github.com/rezwanul-haque/History-Service/src/controllers/history"
	"github.com/rezwanul-haque/History-Service/src/controllers/ping"
	"github.com/rezwanul-haque/History-Service/src/utils/consts"
)

const (
	APIBase = "api/" + consts.APIVersion
)

func mapUrls() {
	router.GET(APIBase+"/ping", ping.Ping)

	router.GET(APIBase+"/history", history.Get)
}
