package usecase

import (
	"context"
)

type (
	Rle interface {
		RunLengthEncode(context.Context, []string) ([]string, error)
		RunLengthDecode(context.Context, []string) ([]string, error)
	}

	RleRepo interface {
		RunLengthEncode(context.Context, []string) ([]string, error)
		RunLengthDecode(context.Context, []string) ([]string, error)
	}
)
