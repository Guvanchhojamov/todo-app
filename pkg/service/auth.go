package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/guvanchhojamov/app-todo/pkg/repository"
	"time"
)

const (
	salt      = "qwerAhjks4dsa"
	tokenTLL  = time.Hour * 12
	signedKey = "165dsaASD#KO48946"
)

type AuthService struct {
	repo repository.Authorization
}

type customTokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	return as.repo.CreateUser(user)
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))) //Sum()  is hashedString+salt
}

// Sign in msethods

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from DB
	user, err := as.repo.GetUserFromDB(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customTokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenTLL)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		user.Id,
	})
	return token.SignedString([]byte(signedKey))
}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &customTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signin method on Parse!")
		}
		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*customTokenClaims)
	if !ok {
		return 0, errors.New("token Claims are not of type *interface: *customTokenClaims")
	}
	return claims.UserId, nil
}
