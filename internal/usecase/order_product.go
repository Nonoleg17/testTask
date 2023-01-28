package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"testCase/internal/entity"
)

type OrderProductUseCase struct {
	repo OrderProductRepo
}

func NewOrderProductUseCase(r OrderProductRepo) *OrderProductUseCase {
	return &OrderProductUseCase{r}
}

func (uc *OrderProductUseCase) CreateOrderProduct(ctx context.Context, o *entity.OrderProduct) error {
	return uc.repo.CreateOrderProduct(ctx, o)
}

func (uc *OrderProductUseCase) ClearAllOrderProduct(ctx context.Context, orderId uuid.UUID) error {
	return uc.repo.ClearAllOrderProduct(ctx, orderId)
}
func (uc *OrderProductUseCase) DeleteOrderProduct(ctx context.Context) error {
	return uc.repo.DeleteOrderProduct(ctx)
}

func (uc *OrderProductUseCase) UpdateOrderProduct(ctx context.Context) error {
	return uc.repo.UpdateOrderProduct(ctx)
}
