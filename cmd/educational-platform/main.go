package main

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/config"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/users_repo"
	dbclients "github.com/deevins/educational-platform-backend/pkg/db/clients"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development" // Default to development if ENVIRONMENT is not set
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatalf("Error loading .env.%s file: %v", env, err)
	}

	if err := initConfig(); err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	//db, err := dbclients.NewPostgres(ctx, config.Config{
	//	Host:     os.Getenv("DB_HOST"),
	//	Port:     os.Getenv("DB_PORT"),
	//	User:     os.Getenv("DB_USER"),
	//	Password: os.Getenv("DB_PASSWORD"),
	//	Dbname:   os.Getenv("DB_DBNAME"),
	//	SSLMode:  os.Getenv("DB_SSL"),
	//}) // returned struct get interfaces
	//if err != nil {
	//	log.Fatalf("can not connect to db method clients.NewPgDB: %s", err)
	//}

	db, err := dbclients.NewPostgres(ctx, config.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Dbname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}) // returned struct get interfaces
	if err != nil {
		log.Fatalf("can not connect to db method clients.NewPgDB: %s", err)
	}

	// Создаем новый экземпляр роутера Gin
	router := gin.Default()

	// Определяем обработчик для GET запросов на путь "/"
	router.GET("/api/v1/courses/", func(c *gin.Context) {
		// Возвращаем текст "Hello, World!" с кодом состояния 200
		c.String(http.StatusOK, "Hello, World!")
	})

	// Запускаем сервер на порту 8080

	if err = router.Run(":8080"); err != nil {
		log.Fatalf("can not run server: %s", err.Error())
	}

	defer db.GetPool(ctx).Close()

}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func testFunc(ctx *gin.Context) {
	var input users_repo.HumanResourcesUser
	if err := ctx.BindJSON(&input); err != nil {
		return
	}

	id, err := users_repo.GetUsers(ctx, input)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
