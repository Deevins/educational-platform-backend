package main

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/config"
	"github.com/deevins/educational-platform-backend/internal/handler"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/S3"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/courses"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/directories"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/users"
	"github.com/deevins/educational-platform-backend/internal/servers"
	"github.com/deevins/educational-platform-backend/internal/service/auth"
	"github.com/deevins/educational-platform-backend/internal/service/course"
	"github.com/deevins/educational-platform-backend/internal/service/directory"
	"github.com/deevins/educational-platform-backend/internal/service/user"
	dbclients "github.com/deevins/educational-platform-backend/pkg/db/clients"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	db, err := dbclients.NewPostgres(ctx, config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSL"),
	})
	if err != nil {
		log.Fatalf("can not connect to db with err: %s", err)
	}

	// Инициализация соединения с Minio
	minioClient := S3.NewMinioClient()
	err = minioClient.InitMinio()
	if err != nil {
		log.Fatalf("Ошибка инициализации Minio: %v", err)
	}

	userRepo := users.New(db)
	courseRepo := courses.New(db)
	directoryRepo := directories.New(db)
	authSvc := auth.NewService(userRepo)
	userSvc := user.NewService(userRepo, minioClient)
	courseSvc := course.NewService(courseRepo, minioClient)
	directorySvc := directory.NewService(directoryRepo)

	handlers := handler.NewHandler(authSvc, userSvc, courseSvc, nil, directorySvc)

	srv := new(servers.HTTPServer)
	go func() {
		if err := srv.Run(viper.GetString("http_server.port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Println("Application is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error on server while shutting down: [%s]", err)
	}

	defer db.GetPool(ctx).Close()

	//grpcServerPort := viper.GetString("grpc_server.port")
	//lsn, err := net.Listen("tcp", grpcServerPort)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//grpcServer := grpc.NewServer()
	//
	//pb.RegisterUserServiceV1Server(grpcServer, NewUserImplementation(userService))

	//go func() {
	//	log.Printf("gRPC server successfully started on port %s", lsn.Addr().String())
	//	if err := grpcServer.Serve(lsn); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

//ENV=development
//DB_HOST=localhost
//DB_PORT=5432
//DB_USER=Shili
//DB_PASSWORD=postgres
//DB_DBNAME=pg
//DB_SSL=disable
//SIGNING_KEY=sdfasddsfsdfsdfsdfsd
//SALT=sdfsdfrdff$#@Q#$sdfsdfd12312sadsdfsdf
//
//
//MINIO_ENDPOINT=localhost:9000
//MINIO_ROOT_USER=root
//MINIO_ROOT_PASSWORD=zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG
//MINIO_BUCKET_NAME=test-bucket
//MINIO_USE_SSL=false
//FILE_TIME_EXPIRATION=24 # в часах
