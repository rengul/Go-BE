package consumption

import (
	"context"
	"re-home/models"
	mod "re-home/models"
)

type Repository interface {
	Get(ctx context.Context, user *models.User) ([]*mod.Consumption, error)
}
