package main

import (
	"flag"
	"fmt"
	"os"
	"update-api/pkg/api"
	"update-api/pkg/configs"
	"update-api/pkg/databases"
	"update-api/pkg/logging"
	"update-api/pkg/services"

	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "path to config.yaml file")
	flag.Parse()

	config, err := configs.NewConfig(*configPath)
	if err != nil {
		fmt.Printf("failed read config: %s", err.Error())
		os.Exit(1)
	}

	log := logging.InitLogger(config.LogLevel)

	updateDB := databases.NewMySqlUpdateDatabase()
	if err := updateDB.Connect(config.DSN); err != nil {
		log.Error("Failed connect to db", zap.Error(err))
		os.Exit(1)
	}

	if err := api.NewServerAPI(services.NewDefaultServiceCollector(updateDB)).
		Start(fmt.Sprintf("%s:%s", config.Addresss, config.Port)); err != nil {
		log.Error("Failed start server", zap.Error(err))
		os.Exit(1)
	}

}
