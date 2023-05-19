package repository

import "gin_test/crud_format_template/model"

type TagsRepository interface {
	Save(tags model.Tags) error
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() []model.Tags
}
