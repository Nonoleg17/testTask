package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"testCase/internal/entity"
)

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{r}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, u *entity.User) error {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, u *entity.User) error {
	return uc.repo.UpdateUser(ctx, u)
}
