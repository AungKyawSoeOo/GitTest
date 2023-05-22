package main

import (
	"gin_test/bulletin_board/config"
	"gin_test/bulletin_board/controller"
	"gin_test/bulletin_board/helper"
	"gin_test/bulletin_board/interfaces"
	"gin_test/bulletin_board/model"
	"gin_test/bulletin_board/router"
	"gin_test/bulletin_board/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server Started")

	db := config.ConnectDatabase()
	validate := validator.New()
	db.Table("posts").AutoMigrate(&model.Posts{})
	db.Table("users").AutoMigrate(&model.User{})

	// Users
	userInterface := interfaces.NewUsersInterfaceImpl(db)
	authService := service.NewAuthServiceImpl(userInterface, validate)
	authController := controller.NewAuthController(authService)
	userController := controller.NewUsercontroller(userInterface)

	// Posts
	postsInterface := interfaces.NewPostsRepositoryImpl(db)
	postsService := service.NewPostsRepositoryImpl(postsInterface, validate)
	postsController := controller.NewPostsController(postsService)

	routes := router.NewRouter(authController, userController, postsController, userInterface)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
