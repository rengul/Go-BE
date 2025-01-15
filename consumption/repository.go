package consumption

import (
	"context"
	mod "re-home/models"
)

type Repository interface {
	Get(ctx context.Context, userId string) ([]*mod.Consumption, error)
}
