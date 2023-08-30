package main

import (
	app "avito-dynamic-segment-back"
	"avito-dynamic-segment-back/pkg/handler"
	"avito-dynamic-segment-back/pkg/repository"
	"avito-dynamic-segment-back/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	password := ""
	if err := godotenv.Load(); err != nil {
		password = "password"
		log.Fatalf("error not have .env file or password parameter there: %s", err.Error())
	} else {
		password = os.Getenv("DB_PASSWORD")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: password,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(app.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
