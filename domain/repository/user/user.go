package user

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

type User interface {
	Create(context.Context, *entity.User) (int64, error)
	Update(context.Context, *entity.User) error
	Delete(context.Context, int64) error
	Detail(context.Context, int64) (*entity.User, error)
	List(context.Context, uint16) ([]*entity.User, error)
}
