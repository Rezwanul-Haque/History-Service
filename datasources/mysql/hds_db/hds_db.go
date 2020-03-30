package hds_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rezwanul-haque/History-Service/logger"
	"github.com/rezwanul-haque/History-Service/utils/helpers"
)

var (
	Client *sql.DB

	username = helpers.GoDotEnvVariable("MYSQL_HDS_USERNAME")
	password = helpers.GoDotEnvVariable("MYSQL_HDS_PASSWORD")
	host     = helpers.GoDotEnvVariable("MYSQL_HDS_HOST")
	port     = helpers.GoDotEnvVariable("MYSQL_HDS_PORT")
	schema   = helpers.GoDotEnvVariable("MYSQL_HDS_SCHEMA")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, host, port, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	logger.Info("database successfully configured")
}
