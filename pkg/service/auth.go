package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	todo "learn-rest-api.go"
	"learn-rest-api.go/pkg/repository"
)

const (
	salt       = "hui"
	tokenTTL   = 12 * time.Hour
	signingKey = "dfjkhdksdjhaffhsadkjf"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (service *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return service.repo.CreateUser(user)
}

func (service *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := service.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		fmt.Println("Error while getting user from db")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	fmt.Println(token)

	return token.SignedString([]byte(signingKey))
}

func (service *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Not correct type of JWT token")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
