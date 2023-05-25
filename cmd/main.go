package main

import (
	"os"

	lavka "github.com/EltIsma/YandexLavka"
	"github.com/EltIsma/YandexLavka/pkg/handler"
	"github.com/EltIsma/YandexLavka/pkg/repository"
	"github.com/EltIsma/YandexLavka/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil{
		logrus.Fatalf("error: %s", err.Error())
	}

	if err:= godotenv.Load(); err != nil{
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewSrvice(repos)
	handlers := handler.NewHandler(services)
	srv := new(lavka.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()) ; err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
 
 }