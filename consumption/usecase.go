package consumption

import (
	"context"
	"re-home/models"
)

type UseCase interface {
	GetConsumption(ctx context.Context, user *models.User) error
}
