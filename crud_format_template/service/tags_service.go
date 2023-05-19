package service

import (
	"gin_test/crud_format_template/data/request"
	"gin_test/crud_format_template/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest) error
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
