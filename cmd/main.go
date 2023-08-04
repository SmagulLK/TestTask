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
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Sslmode:  viper.GetString("db.sslmode"),
		Username: viper.GetString("db.username"),
		DbName:   viper.GetString("db.dbname"),
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
