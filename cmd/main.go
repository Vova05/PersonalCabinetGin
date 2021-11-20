package main

import (
	"PersonalCabinetGin"
	"PersonalCabinetGin/pkg/handler"
	"PersonalCabinetGin/pkg/repository"
	"PersonalCabinetGin/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil{
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil{
		log.Fatalf("error initializing env: %s", err.Error())
	}
	db, errDB := repository.NewMySQLDB(repository.Config{
		MdbUser: viper.GetString("db.dbUser"),
		MdbPass: os.Getenv("DB_PASSWORD"),
		MdbHost: viper.GetString("db.dbHost"),
		MdbName: viper.GetString("db.dbName"),
	})
	if errDB != nil{
		log.Fatalf("failed to initialize db: %s", errDB.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(PersonalCabinetGin.Server)
	if err := srv.Run(viper.GetString("port"),handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error{
	viper.AddConfigPath("cmd/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}