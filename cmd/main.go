package main

import (
	"TestTask/handler"
	"TestTask/repository"
	"TestTask/server"
	"TestTask/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewDB(repository.Config{
		Host:     os.Getenv("db_host"),
		Port:     os.Getenv("db_port"),
		Sslmode:  os.Getenv("db_sslmode"),
		Username: os.Getenv("db_username"),
		DbName:   os.Getenv("db_dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	//db, err := repository.NewDB(repository.Config{
	//	Host:     "containers-us-west-133.railway.app",
	//	Port:     "7436",
	//	Sslmode:  "require",
	//	Username: "postgres",
	//	DbName:   "railway",
	//	Password: "kdIxEVV3CYBY0nQTCxjm",
	//})
	if err != nil {
		log.Fatalf("Error of connecting to DB: %s \nx", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	serv := new(server.Server)
	if err := serv.Run("8000", handlers.InitRoutes()); err != nil {
		panic(err)
	}

}
