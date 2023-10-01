package main

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/config"
	dbclients "github.com/deevins/educational-platform-backend/pkg/db/clients"
	"github.com/spf13/viper"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := initConfig(); err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	db, err := dbclients.NewPGClient(ctx, config.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Dbname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("can not connect to db method clients.NewPgDB: %s", err)
	}

	defer db.GetPool(ctx).Close()

}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
