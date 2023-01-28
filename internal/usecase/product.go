package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"testCase/internal/entity"
)

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(r ProductRepo) *ProductUseCase {
	return &ProductUseCase{r}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, u *entity.Product) error {
	return uc.repo.CreateProduct(ctx, u)
}

func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteProduct(ctx, id)
}

func (uc *ProductUseCase) UpdateProduct(ctx context.Context, u *entity.Product) error {
	return uc.repo.UpdateProduct(ctx, u)
}
