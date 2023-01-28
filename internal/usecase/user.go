package usecase

import (
	"context"
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

func (uc *UserUseCase) DeleteUser(ctx context.Context, u entity.User) error {
	return uc.repo.DeleteUser(ctx, u)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, u entity.User) error {
	return uc.repo.UpdateUser(ctx, u)
}
