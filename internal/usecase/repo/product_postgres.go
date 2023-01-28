package repo

import (
	"context"
	uuid "github.com/satori/go.uuid"
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

func (pr *ProductRepo) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := pr.pg.DbConnect.Delete(&entity.Product{ID: id}).Error; err != nil {
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
