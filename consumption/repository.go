package consumption

import (
	"context"
	"re-home/models"
	mod "re-home/models"
)

type Repository interface {
	GetHeating(ctx context.Context, userId string) ([]*mod.Consumption, error)
	GetHotWater(ctx context.Context, userId string) ([]*mod.Consumption, error)
	GetColdWater(ctx context.Context, userId string) ([]*mod.Consumption, error)
	GetAllConsumption(ctx context.Context, userId string, filter models.Filter) ([]*mod.Consumption, error)
}
