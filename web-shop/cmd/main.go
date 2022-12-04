package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/evgenchikk/itmo-web-shop/config"
	"github.com/evgenchikk/itmo-web-shop/pkg/handler"
	"github.com/evgenchikk/itmo-web-shop/pkg/repository"
	"github.com/evgenchikk/itmo-web-shop/pkg/service"
	"github.com/evgenchikk/itmo-web-shop/server"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Starting with deafult params")
	}
}

func main() {
	ctx := context.Background()
	config := config.NewConfig()

	db, err := repository.NewPostgresDB(
		ctx,
		config.DB.DB_HOST,
		config.DB.DB_PORT,
		config.DB.DB_NAME,
		config.DB.DB_USER,
		config.DB.DB_PASSWD,
	)

	if err != nil {
		log.Fatalf("Database connect failed: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	server := new(server.Server)

	go func() {
		if err := server.Run(config.Port, handler.InitRoutes()); err != nil {
			log.Fatalf("Failed to run serevr: %s", err.Error())
		}
	}()

	log.Printf("Server successfuly started on %s:%s", config.Host, config.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server is shutting down...")

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(ctx); err != nil {
		log.Printf("Error occured on db connection close: %s", err.Error())
	}
}
