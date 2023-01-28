package usecase

import (
	"context"
	"testCase/internal/entity"
)

type (
	User interface {
		CreateUser(context.Context, *entity.User) error
		DeleteUser(context.Context, entity.User) error
		UpdateUser(context.Context, entity.User) error
	}
	Product interface {
		CreateProduct(ctx context.Context) error
	}
	UserRepo interface {
		CreateUser(context.Context, *entity.User) error
		DeleteUser(context.Context, entity.User) error
		UpdateUser(context.Context, entity.User) error
	}
)
