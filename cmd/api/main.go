package main

import (
	"go-digilib/api"
	"go-digilib/api/middlewares"
	"go-digilib/db/drivers"
	"go-digilib/pkg/constant"
	"go-digilib/pkg/fileupload"
	"go-digilib/pkg/utils"
	"log"
	"strconv"
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

	expireDuration, err := strconv.Atoi(utils.GetConfig("JWT_EXPIRE_DURATION"))

	if err != nil {
		log.Fatalf("error when parsing expire duration: %v\n", err)
	}

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: expireDuration,
	}

	var (
		repository = dbConfig.InitDB()
		cloudinary = cloudinaryConfig.InitCloudinary()
		e          = api.NewEcho(repository, cloudinary, jwtConfig)
	)

	drivers.MigrateDB(repository)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
