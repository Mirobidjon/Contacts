package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Mirobidjon/contact-list"
	"github.com/Mirobidjon/contact-list/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	salt = "aSWdkh6465a4dEWdyKHJS"
	timeToken = 12 * time.Hour
	signingKey = "asd12aHJGJHG4sad"
)

// AuthService ...
type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	UserID int
	jwt.StandardClaims
}

// NewAuthService ...
func NewAuthService(c repository.Authorization) *AuthService {
	return &AuthService{
		repo: c,
	}
}

// CreateUser ...
func (s *AuthService) CreateUser(user contact.User) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken ...
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, hashPassword(password))
	if err != nil {
		return "", err
	}

	tk := tokenClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeToken).Unix(),
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	return token.SignedString([]byte(signingKey))
}

// ParseToken token parse and returen user id
func (s *AuthService) ParseToken(token string) (int, error) {
	tk, err := jwt.ParseWithClaims(token, &tokenClaims{}, func (t *jwt.Token) (interface{}, error){
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, nil
	}

	cl, ok := tk.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return cl.UserID, nil
}

func hashPassword(pass string) string{
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}