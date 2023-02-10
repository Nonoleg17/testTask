package usecase

import "context"

type RleUseCase struct {
	repo RleRepo
}

func NewRleUseCase(r RleRepo) *RleUseCase {
	return &RleUseCase{
		repo: r,
	}
}

func (ru *RleUseCase) RunLengthEncode(ctx context.Context, encodeStr []string) ([]string, error) {
	return ru.repo.RunLengthEncode(ctx, encodeStr)
}

func (ru *RleUseCase) RunLengthDecode(ctx context.Context, decodeStr []string) ([]string, error) {
	return ru.repo.RunLengthDecode(ctx, decodeStr)
}
