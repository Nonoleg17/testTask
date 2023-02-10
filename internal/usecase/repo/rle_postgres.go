package repo

import (
	"context"
	"testCase/pkg/postgres"
)

type RleRepo struct {
	pg *postgres.Postgres
}

func NewRleRepo(pg *postgres.Postgres) *RleRepo {
	return &RleRepo{
		pg,
	}

}

func (rr *RleRepo) RunLengthEncode(ctx context.Context, encodeStr []string) ([]string, error) {
	res := make([]string, len(encodeStr))
	for index, value := range encodeStr {
		encodeStr := encode(value)
		res[index] = encodeStr
	}

	return res, nil
}

func (rr *RleRepo) RunLengthDecode(ctx context.Context, decodeStr []string) ([]string, error) {
	res := make([]string, len(decodeStr))
	for index, value := range decodeStr {
		decodeStr := decode(value)
		res[index] = decodeStr
	}

	return res, nil
}
