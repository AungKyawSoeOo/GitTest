package service

import (
	"gin_test/crud_format_template/data/request"
	"gin_test/crud_format_template/data/response"
	"gin_test/crud_format_template/helper"
	"gin_test/crud_format_template/model"
	"gin_test/crud_format_template/repository"

	"github.com/go-playground/validator/v10"
)

func NewTagsRepositoryImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		validate:       validate,
	}
}

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

// Create implements TagsService
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) error {
	err := t.validate.Struct(tags)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			return validationErrs
		}
		return err
	}

	tagModel := model.Tags{
		Name: tags.Name,
	}
	err = t.TagsRepository.Save(tagModel)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements TagsService
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService
func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagsService
func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)
	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements TagsService
func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}
