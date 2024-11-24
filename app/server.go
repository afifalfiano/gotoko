package app

import (
	"flag"
	"log"
	"os"

	"github.com/afifalfiano/gotoko/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoToko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")
	appConfig.AppURL = getEnv("APP_URL", "http://localhost:9000")

	dbConfig.DB_HOST = getEnv("DB_HOST", "localhost")
	dbConfig.DB_USER = getEnv("DB_USER", "root")
	dbConfig.DB_PASSWORD = getEnv("DB_PASSWORD", "")
	dbConfig.DB_NAME = getEnv("DB_NAME", "gotokodb")
	dbConfig.DB_PORT = getEnv("DB_PORT", "3306")
	dbConfig.DB_DRIVER = getEnv("DB_DRIVER", "mysql")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitCommands(dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}
