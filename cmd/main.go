package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	todoapp "github.com/ciiska5/todo-app"
	"github.com/ciiska5/todo-app/pkg/handler"
	"github.com/ciiska5/todo-app/pkg/repository"
	"github.com/ciiska5/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	//задаем логеру JSON формат
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//инициализация БД с необходимыми значениями
	db, err := repository.NewPostgesDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initiaize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Todo-App started")

	quit := make(chan os.Signal, 1)                      // блокировка функции main с помощью канала os.Signal
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // запись в канал os.Signal, когда процесс, в котором выполняется данное приложение, получит от системы сигнал SIGTERM или SIGINT
	<-quit                                               // чтение из канала os.Signal

	logrus.Print("Todo-App is shutting down")

	//остановка сервера
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured while sutting down the server: %s", err.Error())
	}

	//закрытие всех соединений с БД
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured while closing database connections: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
