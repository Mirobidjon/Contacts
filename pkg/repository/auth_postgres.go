package repository

import (
	"github.com/Mirobidjon/contact-list"
	"github.com/jmoiron/sqlx"
)

// AuthPostgres ...
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres ...
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

// CreateUser ...
func (r *AuthPostgres) CreateUser(user contact.User) (int, error){
	var id int
	query := "INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id"

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// GetUser get database user
func (r *AuthPostgres) 	GetUser(username, password string) (contact.User, error) {
	var user contact.User
	query := "SELECT id FROM users WHERE username=$1 AND password=$2"
	err := r.db.Get(&user, query, username, password)
	
	return user, err
}