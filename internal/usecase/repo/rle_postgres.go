package repo

import (
	"context"
	"errors"
	"testCase/internal/entity"
	"testCase/pkg/postgres"
)

type RleRepo struct {
	pg *postgres.Postgres
}
type RleChan struct {
	index int
	value string
}

func NewRleRepo(pg *postgres.Postgres) *RleRepo {
	return &RleRepo{
		pg,
	}

}

func (rr *RleRepo) RunLengthEncode(ctx context.Context, encodeStr []string) ([]string, error) {
	if len(encodeStr) == 0 {
		return nil, errors.New("empty array")
	}
	res := make([]string, len(encodeStr))
	c := make(chan RleChan)
	defer close(c)
	for index, value := range encodeStr {
		go encode(value, index, c)
	}
	for i := 0; i < len(res); i++ {
		encodeData := <-c
		res[encodeData.index] = encodeData.value
	}
	//БД добавлена, если потребуется дальше как-то с этим работать
	for i := 0; i < len(res); i++ {
		if err := rr.pg.DbConnect.Create(&entity.Rle{
			FullString:  encodeStr[i],
			ShortString: res[i],
		}).Error; err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (rr *RleRepo) RunLengthDecode(ctx context.Context, decodeStr []string) ([]string, error) {
	if len(decodeStr) == 0 {
		return nil, errors.New("empty array")
	}
	res := make([]string, len(decodeStr))
	c := make(chan RleChan)
	defer close(c)
	for index, value := range decodeStr {
		go decode(value, index, c)
	}
	for i := 0; i < len(res); i++ {
		decodeData := <-c
		res[decodeData.index] = decodeData.value
	}
	//БД добавлена, если потребуется дальше как-то с этим работать
	for i := 0; i < len(res); i++ {
		if err := rr.pg.DbConnect.Create(&entity.Rle{
			FullString:  res[i],
			ShortString: decodeStr[i],
		}).Error; err != nil {
			return nil, err
		}
	}
	return res, nil
}
