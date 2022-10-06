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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	var (
		isTest bool
		repo   *repository.Repository
	)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("err loading env variables %s", err.Error())
	}

	if os.Getenv("environment") == "test" {
		isTest = true
	}

	if isTest {
		repo = repository.NewRepository(nil, isTest)
	} else {
		db, err := repository.NewPostgresDB(repository.Config{
			SSLmode:  os.Getenv("sslmode"),
			Host:     os.Getenv("host"),
			Port:     os.Getenv("port"),
			Username: os.Getenv("username"),
			Password: os.Getenv("password"),
			DBname:   os.Getenv("dbname"),
		})

		if err != nil {
			log.Fatalf("failed initialized db: %s", err.Error())
		}

		repo = repository.NewRepository(db, isTest)
	}

	port := os.Getenv("HTTP_PORT")
	log.Printf("server started on port %s", port)

	service := service.NewService(repo)
	handlers := handler.NewHandler(service)

	srv := new(contact.Server)

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("err occured while httpserver running : %s", err.Error())
	}
}
