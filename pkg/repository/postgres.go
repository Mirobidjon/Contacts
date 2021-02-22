package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Config ...
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLmode  string
}

// NewPostgresDB ...
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLmode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected database")
	return db, err
}
