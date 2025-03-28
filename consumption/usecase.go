package consumption

import (
	"context"
	"re-home/models"
)

type UseCase interface {
	GetConsumption(ctx context.Context, userId string, filter models.Filter) error
}
