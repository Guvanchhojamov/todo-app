package main

import (
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/guvanchhojamov/app-todo/pkg/handler"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/guvanchhojamov/app-todo/pkg/repository"
	"github.com/guvanchhojamov/app-todo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := inintConfig(); err != nil {
		logrus.Fatal("Init Configs Error", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error Loading env variables ", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Configs{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("Failed to initialize db ", err.Error())
	}
	// Adding Dependecy injections    HANDLER->SERVICE->REPOSITORY
	repos := repository.NewRepository(db) // *sqlx.DB
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	//

	srv := new(model.Server)
	go func() {
		if err := srv.Run(viper.GetString("local_port"), handlers.InitRoutes()); err != nil {
			logrus.Fatal("Server Running Error", err.Error())
		}
	}()
	logrus.Println("Server is runnning...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err = srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("shut down app error: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("close db connect err: %s", err.Error())
	}

}

func inintConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
