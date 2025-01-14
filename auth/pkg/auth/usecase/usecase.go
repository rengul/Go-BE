package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"re-home/auth/pkg/auth"
	"re-home/models"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

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

	user, err := a.repo.Get(ctx, user.Username, user.Password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		User: user,
	})

	return token.SignedString(a.signingKey)
}

func (a *Authorizer) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
