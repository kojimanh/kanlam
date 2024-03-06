package config

import (
	"log"
	"os"

	constant "kanlam/const"
	entity "kanlam/entity/config.entity"

	"github.com/joho/godotenv"
)

func SetupConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetServerConfig() entity.ServerConfig {
	serverHost := os.Getenv(constant.ConfigKeyMap.ServerHost)

	serverConfig := entity.ServerConfig{
		Server: serverHost,
	}

	return serverConfig
}

func GetDbConfig() entity.DbConfig {
	dbConfigUrl := os.Getenv(constant.ConfigKeyMap.DbConfigUrl)

	return entity.DbConfig{
		Url: dbConfigUrl,
	}
}

func GetJwtConfig() entity.JwtConfig {
	secretConfig := os.Getenv(constant.ConfigKeyMap.JwtSecret)

	return entity.JwtConfig{
		Secret: secretConfig,
	}
}
