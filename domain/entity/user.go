package entity

import "strings"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func NewUser(name string) *User {
	names := strings.Split(name, " ")
	return &User{
		Name:     names[0],
		Lastname: names[1],
	}
}
