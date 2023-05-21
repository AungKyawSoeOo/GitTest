package controller

import (
	"gin_test/crud_format_template/helper"
	"gin_test/crud_format_template/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersInterface repository.UsersInterface
}

func NewUsercontroller(interfaces repository.UsersInterface) *UsersController {
	return &UsersController{UsersInterface: interfaces}
}

func (controller *UsersController) GetUsers(ctx *gin.Context) {
	users := controller.UsersInterface.FindAll()
	helper.ResponseHandler(ctx, http.StatusOK, "Get All Users Success.", users)
}
