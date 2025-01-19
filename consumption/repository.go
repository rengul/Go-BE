package consumption

import (
	"context"
	mod "re-home/models"
)

type Repository interface {
	GetHeating(ctx context.Context, userId string) ([]*mod.Consumption, error)
	GetHotWater(ctx context.Context, userId string) ([]*mod.Consumption, error)
	GetColdWater(ctx context.Context, userId string) ([]*mod.Consumption, error)
	GetAllConsumption(ctx context.Context, userId string) ([]*mod.Consumption, error)
}
