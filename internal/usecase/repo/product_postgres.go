package repo

import (
	"context"
	"testCase/pkg/postgres"
)

type ProductRepo struct {
	*postgres.Postgres
}

func NewProductRepo(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{
		pg,
	}

}

func (ur *UserRepo) CreateProduct(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) DeleteProduct(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) UpdateProduct(ctx context.Context) error {
	return nil
}
