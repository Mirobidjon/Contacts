package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Mirobidjon/contact-list"
	"github.com/Mirobidjon/contact-list/pkg/handler"
	"github.com/Mirobidjon/contact-list/pkg/repository"
	"github.com/Mirobidjon/contact-list/pkg/service"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("err loading env variables %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		SSLmode: os.Getenv("sslmode"),
		Host: os.Getenv("host"),
		Port: os.Getenv("port"),
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
		DBname: os.Getenv("dbname"),
	})

	if err != nil {
		log.Fatalf("failed initialized db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handlers := handler.NewHandler(service)


	srv := new(contact.Server)

	if err := srv.Run("4000", handlers.InitRoutes()); err != nil {
		log.Fatalf("err occured while httpserver running : %s", err.Error())
	}
}
