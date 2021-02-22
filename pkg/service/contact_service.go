package service

import (
	"errors"

	"github.com/Mirobidjon/contact-list"
	"github.com/Mirobidjon/contact-list/pkg/repository"
)

// ContactService sevice contacts
type ContactService struct {
	repo repository.Contacts
}

// NewContactService ... 
func NewContactService(repo repository.Contacts) *ContactService {
	return &ContactService{repo}
}

// Create new Contacts
func (s *ContactService) Create(name, phone string, userID int) (int, error) {
	return s.repo.Create(name, phone, userID)
}

// GetAll contacts 
func (s *ContactService) GetAll(userID int) ([]contact.Contact, error) {
	return s.repo.GetAll(userID)
}

// GetByID users contact
func (s *ContactService) GetByID(userID, contactID int) (contact.Contact, error) {
	return s.repo.GetByID(userID, contactID)
}

// Delete contact
func (s *ContactService) Delete(userID, contactID int) error {
	return s.repo.Delete(userID, contactID)
}

// Update contact services
func (s *ContactService) Update(userID, contactID int, input contact.DefaultContact) error {
	if input.Name == "" && input.Phone == "" {
		return errors.New("contact hasn't values. ")
	}
	
	return s.repo.Update(userID,contactID, input)
}