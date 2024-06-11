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
	Create(context.Context, *entity.Address) (int64, error)
	Update(context.Context, *entity.Address) error
	Delete(context.Context, int64) error
	Detail(context.Context, int64) (*entity.Address, error)
	List(context.Context, uint16) ([]*entity.Address, error)
}
