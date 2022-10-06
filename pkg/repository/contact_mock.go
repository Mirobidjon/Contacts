package repository

import (
	"errors"
	"sort"

	"github.com/Mirobidjon/contact-list"
)

// ContactMock ...
type ContactMock struct {
	contacts     map[int]contact.Contact
	count        int
	userContacts map[int][]int
}

// NewContactMock ...
func NewContactMock() *ContactMock {
	return &ContactMock{
		contacts:     make(map[int]contact.Contact),
		count:        0,
		userContacts: make(map[int][]int),
	}
}

// Create new Contact
func (r *ContactMock) Create(name, phone string, userID int) (int, error) {
	r.count++
	contact := contact.Contact{
		ID:     r.count,
		Name:   name,
		Phone:  phone,
		UserID: userID,
	}

	r.contacts[r.count] = contact
	r.userContacts[userID] = append(r.userContacts[userID], r.count)
	sort.Slice(r.userContacts[userID], func(i, j int) bool {
		return r.userContacts[userID][i] < r.userContacts[userID][j]
	})

	return r.count, nil
}

// GetAll Contacts from database
func (r *ContactMock) GetAll(userID int) ([]contact.Contact, error) {
	var contacts []contact.Contact
	for _, id := range r.userContacts[userID] {
		contacts = append(contacts, r.contacts[id])
	}

	return contacts, nil
}

// GetByID contact from database
func (r *ContactMock) GetByID(userID, contactID int) (contact.Contact, error) {
	contact, ok := r.contacts[contactID]
	if !ok {
		return contact, errors.New("contact not found")
	}

	if contact.UserID != userID {
		return contact, errors.New("contact not found")
	}

	return contact, nil
}

// Delete contact from database
func (r *ContactMock) Delete(userID, contactID int) error {
	_, ok := r.contacts[contactID]
	if !ok {
		return errors.New("contact not found")
	}

	delete(r.contacts, contactID)
	for i, id := range r.userContacts[userID] {
		if id == contactID {
			r.userContacts[userID] = append(r.userContacts[userID][:i], r.userContacts[userID][i+1:]...)
			break
		}
	}

	return nil
}

// Update contact database
func (r *ContactMock) Update(userID, contactID int, input contact.DefaultContact) error {
	contact, ok := r.contacts[contactID]
	if !ok {
		return errors.New("contact not found")
	}

	contact.Name = input.Name
	contact.Phone = input.Phone
	r.contacts[contactID] = contact

	return nil
}
