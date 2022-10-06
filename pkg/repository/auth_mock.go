package repository

import (
	"errors"

	"github.com/Mirobidjon/contact-list"
)

type AuthMock struct {
	count     int
	users     map[int]contact.User
	usernames map[string]int
}

func NewAuthMock() *AuthMock {
	return &AuthMock{
		count:     0,
		users:     make(map[int]contact.User),
		usernames: make(map[string]int),
	}
}

// CreateUser ...
func (r *AuthMock) CreateUser(user contact.User) (int, error) {
	id, ok := r.usernames[user.Username+user.Password]
	if ok {
		return id, nil
	}

	r.count++
	r.users[r.count] = user
	r.usernames[user.Username+user.Password] = r.count

	return r.count, nil
}

// GetUser get database user
func (r *AuthMock) GetUser(username, password string) (contact.User, error) {
	id, ok := r.usernames[username+password]
	if !ok {
		return contact.User{}, errors.New("user not found")
	}

	return r.users[id], nil
}
