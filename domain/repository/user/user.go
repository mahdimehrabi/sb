package user

import (
	"context"
	"errors"
	"m1-article-service/domain/entity"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrNotFound     = errors.New("not found")
)

type User interface {
	Create(context.Context, *entity.User) (int64, error)
	Detail(context.Context, int64) (*entity.User, error)
}
