package repo

import (
	"context"
	"testCase/internal/entity"
	"testCase/pkg/postgres"
)

type ProductRepo struct {
	pg *postgres.Postgres
}

func NewProductRepo(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{
		pg,
	}

}

func (pr *ProductRepo) CreateProduct(ctx context.Context, p *entity.Product) error {
	if err := pr.pg.DbConnect.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepo) DeleteProduct(ctx context.Context, id string) error {
	if err := pr.pg.DbConnect.Delete(&entity.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepo) UpdateProduct(ctx context.Context, p *entity.Product) error {
	if err := pr.pg.DbConnect.Model(p).Update(&entity.Product{
		Description: p.Description,
		Price:       p.Price,
		Currency:    p.Currency,
		LeftInStock: p.LeftInStock,
	}).Error; err != nil {
		return err
	}
	return nil
}
