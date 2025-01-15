package consumption

import (
	"context"
)

type UseCase interface {
	GetConsumption(ctx context.Context, userId string) error
}
