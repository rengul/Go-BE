package usecase

import (
	"context"
	"re-home/consumption"
	"re-home/models"
)

type ConsumptionUseCase struct {
	ConsumptionRepo consumption.Repository
}

func NewConsumptionUseCase(ConsumptionRepo consumption.Repository) *ConsumptionUseCase {
	return &ConsumptionUseCase{
		ConsumptionRepo: ConsumptionRepo,
	}
}

func (c ConsumptionUseCase) GetConsumption(ctx context.Context, userId string, action models.Action) ([]*models.Consumption, error) {
	if action == models.Heating {
		return c.ConsumptionRepo.GetHeating(ctx, userId)
	}

	if action == models.HotWater {
		return c.ConsumptionRepo.GetHotWater(ctx, userId)
	}

	if action == models.ColdWater {
		return c.ConsumptionRepo.GetColdWater(ctx, userId)
	}

	return c.ConsumptionRepo.GetAllConsumption(ctx, userId)

}
