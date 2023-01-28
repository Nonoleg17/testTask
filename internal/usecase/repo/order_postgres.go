package repo

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"testCase/internal/entity"
	"testCase/pkg/postgres"
)

type OrderRepo struct {
	pg *postgres.Postgres
}

func NewOrderRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{
		pg,
	}

}

func (or *OrderRepo) CreateOrder(ctx context.Context, o *entity.Order) (uuid.UUID, error) {
	if err := or.pg.DbConnect.Create(o).Error; err != nil {
		return uuid.UUID{}, err
	}

	return o.ID, nil
}

func (or *OrderRepo) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	if err := or.pg.DbConnect.Delete(&entity.Order{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

func (or *OrderRepo) UpdateOrder(ctx context.Context, o *entity.Order) error {
	if err := or.pg.DbConnect.Model(o).Update(&entity.Order{
		UserId: o.UserId,
	}).Error; err != nil {
		return err
	}
	return nil
}
