package repo

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testCase/internal/entity"
	"testCase/pkg/postgres"
)

type OrderProductRepo struct {
	pg *postgres.Postgres
}

func NewOrderProductRepo(pg *postgres.Postgres) *OrderProductRepo {
	return &OrderProductRepo{
		pg,
	}

}
func (or *OrderProductRepo) CreateOrderProduct(ctx context.Context, o *entity.OrderProduct) error {
	product := &entity.Product{ID: o.ProductId}
	if err := or.pg.DbConnect.First(product).Error; err != nil {
		return err
	}
	if product.LeftInStock-o.CountProducts < 0 {
		return errors.New(fmt.Sprintf("not enough %s", product.Description))
	}
	if err := or.pg.DbConnect.Create(o).Error; err != nil {
		return err
	}
	if err := or.pg.DbConnect.Model(product).Update(&entity.Product{
		LeftInStock: product.LeftInStock - o.CountProducts,
	}).Error; err != nil {
		return err
	}
	return nil

}
func (or *OrderProductRepo) ClearAllOrderProduct(ctx context.Context, orderId uuid.UUID) error {
	var orderProduct []entity.OrderProduct
	if err := or.pg.DbConnect.Find(&orderProduct).Where("order_id = ?", orderId).Error; err != nil {
		return err
	}
	for _, value := range orderProduct {
		var product entity.Product
		if err := or.pg.DbConnect.Find(&product).Where("id = ?", value.ProductId).Error; err != nil {
			return err
		}
		if err := or.pg.DbConnect.Model(&entity.Product{}).Update(&entity.Product{
			LeftInStock: value.CountProducts + product.LeftInStock,
		}).Where("id = ?", product.ID).Error; err != nil {
			return err
		}
	}
	if err := or.pg.DbConnect.Delete(&entity.OrderProduct{}, "order_id = ?", orderId).
		Error; err != nil {
		return err
	}
	return nil
}
func (or *OrderProductRepo) DeleteOrderProduct(ctx context.Context) error {
	return nil
}

func (or *OrderProductRepo) UpdateOrderProduct(ctx context.Context) error {
	return nil
}
