package repository

import (
	"fmt"
	"strings"

	"github.com/Mirobidjon/contact-list"
	"github.com/jmoiron/sqlx"
)

// ContactPostgres ... 
type ContactPostgres struct {
	db *sqlx.DB
}

// NewContactPostgres ...
func NewContactPostgres(db *sqlx.DB) *ContactPostgres {
	return &ContactPostgres{db}
}

// Create new Contact
func (r *ContactPostgres) Create(name, phone string, userID int) (int, error){
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	 
	var id int
	query := "INSERT INTO contacts (name, phone, user_id) VALUES ($1, $2, $3) RETURNING id"
	row := tx.QueryRow(query, name, phone, userID)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

// GetAll Contacts from database
func (r *ContactPostgres) GetAll(userID int) ([]contact.Contact, error) {
	var contacts []contact.Contact
	query := "SELECT id, name, phone, user_id FROM contacts WHERE user_id=$1"

	err := r.db.Select(&contacts, query, userID)

	return contacts, err
}

// GetByID contact from database
func (r *ContactPostgres) GetByID(userID, contactID int) (contact.Contact, error) {
	var contact contact.Contact
	query := "SELECT id, name, phone, user_id FROM contacts WHERE user_id=$1 AND id=$2"
	
	err := r.db.Get(&contact, query, userID, contactID)
	
	return contact, err
}

// Delete contact from database
func (r *ContactPostgres) Delete(userID, contactID int) error {
	query := "DELETE FROM contacts WHERE user_id=$1 AND id=$2"

	_, err := r.db.Exec(query, userID, contactID)

	return err
}

// Update contact database
func (r *ContactPostgres) Update(userID, contactID int, input contact.DefaultContact) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != "" {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argID))
		args = append(args, input.Name)
		argID++
	}

	if input.Phone != "" {
		setValue = append(setValue, fmt.Sprintf("phone=$%d", argID))
		args = append(args, input.Phone)
		argID++
	}

	argString := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE contacts SET %s WHERE user_id=$%d AND id=$%d", argString, argID, argID+1)
	args = append(args, userID, contactID)
	
	_, err := r.db.Exec(query, args...)

	return err
}