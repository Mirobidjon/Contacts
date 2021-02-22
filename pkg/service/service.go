package service

import (
	"github.com/Mirobidjon/contact-list"
	"github.com/Mirobidjon/contact-list/pkg/repository"
)

// Authorization ...
type Authorization interface {
	CreateUser(user contact.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// Contacts ...
type Contacts interface {
	Create(name, phone string, userID int) (int, error)
	GetAll(userID int) ([]contact.Contact, error)
	GetByID(userID, contactID int) (contact.Contact, error)
	Delete(userID, contactID int) error
	Update(userID, contactID int, input contact.DefaultContact) error
}

// Service ...
type Service struct {
	Authorization
	Contacts
}

// NewService ...
func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
		Contacts: NewContactService(r.Contacts),
	}
}