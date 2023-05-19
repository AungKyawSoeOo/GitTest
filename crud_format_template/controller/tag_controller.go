package controller

import (
	"fmt"
	"gin_test/crud_format_template/data/request"
	"gin_test/crud_format_template/data/response"
	"gin_test/crud_format_template/helper"
	"gin_test/crud_format_template/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// create controller
func (controller *TagsController) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	createTagsRequest := request.CreateTagsRequest{
		Name: name,
	}
	err := controller.tagsService.Create(createTagsRequest)
	if err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make(map[string]string)
			for _, fieldErr := range validationErr {
				fieldName := fieldErr.Field()
				errorMessage := ""
				switch fieldErr.Tag() {
				case "required":
					errorMessage = "Field is required"
				case "min":
					errorMessage = fmt.Sprintf("Field must be at least %s characters long", fieldErr.Param())
				case "max":
					errorMessage = fmt.Sprintf("Field must not exceed %s characters", fieldErr.Param())
				default:
					errorMessage = "Field validation failed"
				}
				errorMessages[fieldName] = errorMessage

			}
			fmt.Println(errorMessages)
			ctx.HTML(http.StatusBadRequest, "create.html", gin.H{
				"Errors": errorMessages,
			})
			return
		}
	} else {
		ctx.Redirect(http.StatusFound, "/api/tags")
	}
}

// update controller
func (controller *TagsController) Update(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	fmt.Println(id, "ffopposfusfiugpfsogffssf")
	helper.ErrorPanic(err)
	updateTagsRequest := request.UpdateTagsRequest{
		Id: id,
	}

	if err := ctx.ShouldBind(&updateTagsRequest); err != nil {
		ctx.HTML(http.StatusBadRequest, "update.html", gin.H{
			"Tag":    updateTagsRequest,
			"Errors": err.Error(),
		})
		return
	}
	controller.tagsService.Update(updateTagsRequest)
	ctx.Redirect(http.StatusFound, "/api/tags")
}

// delete controller
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// findById controller
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagsResponse := controller.tagsService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// findAll controller
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"tags": tagResponse,
	})
}

func (controller *TagsController) CreateForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create.html", gin.H{})
}

func (controller *TagsController) UpdateForm(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tag := controller.tagsService.FindById(id)
	ctx.HTML(http.StatusOK, "update.html", gin.H{
		"Tag": tag,
	})
}
