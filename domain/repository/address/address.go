package address

import (
	"context"
	"errors"
	"m1-article-service/domain/entity"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrValidation   = errors.New("validation error")
	ErrNotFound     = errors.New("not found")
)

type Address interface {
	BatchCreate(context.Context, []*entity.Address) error
	Detail(context.Context, int64) (*entity.Address, error)
}
