package repo

import (
	"context"
	uuid "github.com/satori/go.uuid"
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
	if err := ur.pg.DbConnect.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := ur.pg.DbConnect.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) UpdateUser(ctx context.Context, u *entity.User) error {
	if err := ur.pg.DbConnect.Model(u).Update(&entity.User{
		Firstname:  u.Firstname,
		Surname:    u.Surname,
		Middlename: u.Middlename,
		Sex:        u.Sex,
		Age:        u.Age,
	}).Error; err != nil {
		return err
	}
	return nil
}
