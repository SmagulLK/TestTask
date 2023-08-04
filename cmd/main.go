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
	if err := InitConfig(); err != nil {
		log.Fatalf("Error of reading config file: %s \n", err.Error())
	}

	db, err := repository.NewDB(repository.Config{
		Host:     os.Getenv("db.host"),
		Port:     os.Getenv("db.port"),
		Sslmode:  os.Getenv("db.sslmode"),
		Username: os.Getenv("db.username"),
		DbName:   os.Getenv("db.dbname"),
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
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
