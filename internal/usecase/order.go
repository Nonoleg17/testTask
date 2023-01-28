package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"testCase/internal/entity"
)

type OrderUseCase struct {
	repo OrderRepo
}

func NewOrderUseCase(r OrderRepo) *OrderUseCase {
	return &OrderUseCase{r}
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, o *entity.Order) (uuid.UUID, error) {
	return uc.repo.CreateOrder(ctx, o)
}

func (uc *OrderUseCase) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteOrder(ctx, id)
}

func (uc *OrderUseCase) UpdateOrder(ctx context.Context, o *entity.Order) error {
	return uc.repo.UpdateOrder(ctx, o)
}
