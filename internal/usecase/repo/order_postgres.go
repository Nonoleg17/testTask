package repo

import (
	"context"
	"testCase/pkg/postgres"
)

type OrderRepo struct {
	*postgres.Postgres
}

func NewOrderRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{
		pg,
	}

}

func (ur *UserRepo) CreateOrder(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) DeleteOrder(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) UpdateOrder(ctx context.Context) error {
	return nil
}
