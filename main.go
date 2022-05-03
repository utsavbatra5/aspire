package main

import (
	"aspire/api"
	"aspire/config"
	"aspire/logger"
	"aspire/persistency/database"
)

func init() {
	// initialize config, database and logger
	config.InitConfig()
	database.InitConnections()
	logger.InitLogger()
}

func main() {
	api.InitAPI()
}
