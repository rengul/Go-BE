package auth

import (
	"re-home/models"

	"github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}
