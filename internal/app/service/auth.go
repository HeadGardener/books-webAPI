package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/HeadGardener/books-webAPI/internal/app/models"
	"github.com/HeadGardener/books-webAPI/internal/app/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "adnboekvnegnrueidkvme"
	tokenTTL   = 12 * time.Hour
	signingKey = "iermimimcsdovtinbbvr"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateToken(inputUser models.UserInput) (string, error) {
	inputUser.Password = generatePasswordHash(inputUser.Password)
	user, err := s.repos.GetUser(inputUser)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
