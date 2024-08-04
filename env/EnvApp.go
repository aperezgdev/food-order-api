package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvApp struct {
	PORT string
	PORT_DB string
	DB_NAME string
	DB_USER string
	DB_PASSWORD string
}

func NewEnvApp() EnvApp {
	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load()
	
		if err != nil {
			panic(err)
		}
	}

	return EnvApp{
		PORT: os.Getenv("PORT"),
		PORT_DB:      os.Getenv("PORT_DB"),
		DB_USER: os.Getenv("DB_USER"),
		DB_NAME:  os.Getenv("DB_NAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}
}