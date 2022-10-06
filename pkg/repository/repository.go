package repository

import (
	"github.com/Mirobidjon/contact-list"
	"github.com/jmoiron/sqlx"
)

// Authorization ...
type Authorization interface {
	CreateUser(user contact.User) (int, error)
	GetUser(username, password string) (contact.User, error)
}

// Contacts ...
type Contacts interface {
	Create(name, phone string, userID int) (int, error)
	GetAll(userID int) ([]contact.Contact, error)
	GetByID(userID, contactID int) (contact.Contact, error)
	Delete(userID, contactID int) error
	Update(userID, contactID int, input contact.DefaultContact) error
}

// Repository ...
type Repository struct {
	Authorization
	Contacts
}

// NewRepository ...
func NewRepository(db *sqlx.DB, isTest bool) *Repository {
	if isTest {
		return &Repository{
			Authorization: NewAuthMock(),
			Contacts:      NewContactMock(),
		}
	}

	return &Repository{
		Authorization: NewAuthPostgres(db),
		Contacts:      NewContactPostgres(db),
	}
}
