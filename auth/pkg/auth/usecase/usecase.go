package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"re-home/auth/pkg/auth"
	"re-home/models"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	ID string `json:"Id"`
}

type Authorizer struct {
	repo auth.Repository

	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthorizer(repo auth.Repository, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		repo:           repo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (a *Authorizer) SignUp(ctx context.Context, user *models.User) error {
	// Create password hash
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	return a.repo.Insert(ctx, user)
}

func (a *Authorizer) SignIn(ctx context.Context, user *models.User) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
	log.Println("Sign in user: ", user.Username)
	user, err := a.repo.Get(ctx, user.Username, user.Password)
	if err != nil {
		return "", err
	}

	claims := AuthClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *Authorizer) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return "", auth.ErrInvalidAccessToken
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.ID, nil
	}

	return "", auth.ErrInvalidAccessToken
}
