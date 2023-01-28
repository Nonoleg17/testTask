package repo

import (
	"context"
	"fmt"
	"testCase/internal/entity"
	"testCase/pkg/postgres"
)

type UserRepo struct {
	pg *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{
		pg,
	}

}

func (ur *UserRepo) CreateUser(ctx context.Context, u *entity.User) error {
	fmt.Println(ur.pg.DbConnect.DB())
	if err := ur.pg.DbConnect.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, u entity.User) error {
	if err := ur.pg.DbConnect.Delete(u).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) UpdateUser(ctx context.Context, u entity.User) error {
	if err := ur.pg.DbConnect.Update(u).Error; err != nil {
		return err
	}
	return nil
}
