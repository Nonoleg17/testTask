package repo

import (
	"context"
	"testCase/pkg/postgres"
)

type OrderProductRepo struct {
	*postgres.Postgres
}

func NewOrderProductRepo(pg *postgres.Postgres) *OrderProductRepo {
	return &OrderProductRepo{
		pg,
	}

}
func (ur *UserRepo) CreateOrderProduct(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) DeleteOrderProduct(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) UpdateOrderProduct(ctx context.Context) error {
	return nil
}
