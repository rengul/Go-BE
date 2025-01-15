package auth

import (
	"context"
	"re-home/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, error)
}
