package main

import (
	"go-digilib/api"
	"go-digilib/db/drivers"
	"go-digilib/shared/constant"
	"go-digilib/shared/utils"
)

func main() {
	dbConfig := drivers.DBConfig{
		Username: utils.GetConfig(constant.DB_USERNAME),
		Password: utils.GetConfig(constant.DB_PASSWORD),
		Database: utils.GetConfig(constant.DB_NAME),
		Host:     utils.GetConfig(constant.DB_HOST),
		Port:     utils.GetConfig(constant.DB_PORT),
	}

	var (
		repository = dbConfig.InitDB()
		e          = api.NewEcho(repository)
	)

	drivers.MigrateDB(repository)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
