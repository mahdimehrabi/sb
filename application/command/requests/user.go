package requests

import "m1-article-service/domain/entity"

type User struct {
	Name      string    `json:"name" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	Addresses Addresses `json:"addresses" `
}

func (u User) ToEntity() *entity.User {
	return &entity.User{
		Name:     u.Name,
		Lastname: u.Lastname,
	}
}
