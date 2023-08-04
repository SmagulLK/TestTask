package main

import (
	"TestTask/handler"
	"TestTask/repository"
	"TestTask/server"
	"TestTask/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	db, err := repository.NewDB(repository.Config{
		Host:     os.Getenv("db_host"),
		Port:     os.Getenv("db_port"),
		Sslmode:  os.Getenv("db_sslmode"),
		Username: os.Getenv("db_username"),
		DbName:   os.Getenv("db_dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Error of connecting to DB: %s \n", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	serv := new(server.Server)
	if err := serv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		panic(err)
	}

}
