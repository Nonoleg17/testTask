package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"testCase/internal/entity"
)

type (
	User interface {
		CreateUser(context.Context, *entity.User) error
		DeleteUser(context.Context, uuid.UUID) error
		UpdateUser(context.Context, *entity.User) error
	}
	Product interface {
		CreateProduct(context.Context, *entity.Product) error
		DeleteProduct(context.Context, uuid.UUID) error
		UpdateProduct(context.Context, *entity.Product) error
	}
	Friendship interface {
		FollowUser(context.Context, uuid.UUID, uuid.UUID) error
	}
	UserRepo interface {
		CreateUser(context.Context, *entity.User) error
		DeleteUser(context.Context, uuid.UUID) error
		UpdateUser(context.Context, *entity.User) error
	}
	ProductRepo interface {
		CreateProduct(context.Context, *entity.Product) error
		DeleteProduct(context.Context, uuid.UUID) error
		UpdateProduct(context.Context, *entity.Product) error
	}
	FriendshipRepo interface {
		FollowUser(context.Context, uuid.UUID, uuid.UUID) error
	}
)
