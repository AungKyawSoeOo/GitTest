package router

import (
	"gin_test/crud_format_template/controller"
	middlewares "gin_test/crud_format_template/middleware"
	"gin_test/crud_format_template/repository"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controller.AuthController, userController *controller.UsersController, tagsController *controller.TagsController, usersInterface repository.UsersInterface) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static/")
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	apiRouter := router.Group("/api")
	AuthRouter(apiRouter, authController)
	UsersRouter(apiRouter, usersInterface, userController)
	TagsRouter(apiRouter, tagsController, usersInterface)
	AuthForm(apiRouter, authController)
	return router
}

// Auth Form Routes
func AuthForm(router *gin.RouterGroup, authController *controller.AuthController) {
	authForm := router.Group("/")
	{
		authForm.GET("/register", authController.RegisterForm)
		authForm.GET("/login", authController.LoginForm)
	}
}

// Auth Router
func AuthRouter(router *gin.RouterGroup, authController *controller.AuthController) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/login", authController.Login)
	}
}

// User Router
func UsersRouter(router *gin.RouterGroup, usersInterface repository.UsersInterface, usersController *controller.UsersController) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("", middlewares.IsAuth(usersInterface), usersController.GetUsers)
	}
}

// TagRouter
func TagsRouter(router *gin.RouterGroup, tagsController *controller.TagsController, userInterface repository.UsersInterface) {
	tagRouter := router.Group("/tags")
	{
		tagRouter.GET("/create", tagsController.CreateForm)
		tagRouter.GET("/update/:tagId", tagsController.UpdateForm)
		tagRouter.GET("", middlewares.IsAuth(userInterface), tagsController.FindAll)
		tagRouter.GET("/:tagId", middlewares.IsAuth(userInterface), tagsController.FindById)
		tagRouter.POST("", middlewares.IsAuth(userInterface), tagsController.Create)
		tagRouter.POST("/:tagId", middlewares.IsAuth(userInterface), tagsController.Update)
		tagRouter.DELETE("/:tagId", middlewares.IsAuth(userInterface), tagsController.Delete)
	}
}
