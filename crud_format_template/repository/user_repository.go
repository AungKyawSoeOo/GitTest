package repository

import "gin_test/crud_format_template/model"

type UsersInterface interface {
	Save(users model.User)
	Update(users model.User)
	Delete(userId int)
	FindById(userId int) (model.User, error)
	FindAll() []model.User
	FindByEmail(email string) (model.User, error)
	FindByUsername(username string) (model.User, error)
}
