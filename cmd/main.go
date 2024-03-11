package main

import (
	"log"

	"github.com/spf13/viper"
	todo "learn-rest-api.go"
	"learn-rest-api.go/pkg/handler"
	"learn-rest-api.go/pkg/repository"
	"learn-rest-api.go/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error while initializing config : %s ", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server, %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
