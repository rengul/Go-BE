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

func (c ConsumptionUseCase) GetConsumption(ctx context.Context, user *models.User) ([]*models.Consumption, error) {
	return c.ConsumptionRepo.Get(ctx, user)
}
