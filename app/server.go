package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/afifalfiano/gotoko/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_DRIVER   string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)
	server.initializeRoutes()
	server.initalizeDB(dbConfig)
}

func (server *Server) initalizeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DB_USER,
		dbConfig.DB_PASSWORD,
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_NAME)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to the database driver")
	}
}

func (server *Server) dbMigrate() {
	for _, model := range RegisterModel() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migration successfully.")
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoToko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DB_HOST = getEnv("DB_HOST", "localhost")
	dbConfig.DB_USER = getEnv("DB_USER", "root")
	dbConfig.DB_PASSWORD = getEnv("DB_PASSWORD", "")
	dbConfig.DB_NAME = getEnv("DB_NAME", "gotokodb")
	dbConfig.DB_PORT = getEnv("DB_PORT", "3306")
	dbConfig.DB_DRIVER = getEnv("DB_DRIVER", "mysql")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
}

func (server *Server) initCommands(dbConfig DBConfig) {
	server.initalizeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DbSeed(server.DB)

				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
