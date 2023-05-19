package router

import (
	"gin_test/crud_format_template/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static/")
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	tagRouter := baseRouter.Group("/tags")
	tagRouter.GET("/create", tagsController.CreateForm)
	tagRouter.GET("/update/:tagId", tagsController.UpdateForm)
	tagRouter.GET("", tagsController.FindAll)
	tagRouter.GET("/:tagId", tagsController.FindById)
	tagRouter.POST("", tagsController.Create)
	tagRouter.PUT("/:tagId", tagsController.Update)
	tagRouter.DELETE("/:tagId", tagsController.Delete)
	return router
}
