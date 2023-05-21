package service

import (
	"gin_test/crud_format_template/data/request"
	"gin_test/crud_format_template/model"
)

type Authservice interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
	FindByEmail(email string) model.User
}
