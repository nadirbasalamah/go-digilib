package main

import (
	"go-digilib/api"
	"go-digilib/db/drivers"
)

func main() {
	dbConfig := drivers.DBConfig{
		DB_USERNAME: "postgres",
		DB_PASSWORD: "mysecretpassword",
		DB_NAME:     "digilib",
		DB_HOST:     "localhost",
		DB_PORT:     "5432",
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
