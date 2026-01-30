package main

import (
	"go-digilib/api"
	"go-digilib/db/drivers"
	"go-digilib/pkg/constant"
	"go-digilib/pkg/fileupload"
	"go-digilib/pkg/utils"
)

func main() {
	dbConfig := drivers.DBConfig{
		Username: utils.GetConfig(constant.DB_USERNAME),
		Password: utils.GetConfig(constant.DB_PASSWORD),
		Database: utils.GetConfig(constant.DB_NAME),
		Host:     utils.GetConfig(constant.DB_HOST),
		Port:     utils.GetConfig(constant.DB_PORT),
	}

	cloudinaryConfig := fileupload.CloudinaryConfig{
		CloudinaryURL: utils.GetConfig("CLOUDINARY_URL"),
	}

	var (
		repository = dbConfig.InitDB()
		cloudinary = cloudinaryConfig.InitCloudinary()
		e          = api.NewEcho(repository, cloudinary)
	)

	drivers.MigrateDB(repository)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
