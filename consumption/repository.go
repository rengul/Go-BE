package consumption

import (
	"context"
	"re-home/models"
	mod "re-home/models"
)

type Repository interface {
	//SignUp(ctx context.Context, user *models.User) error
	//SignIn(ctx context.Context, user *models.User) (string, error)
	Get(ctx context.Context, user *models.User) ([]*mod.Consumption, error)
}
