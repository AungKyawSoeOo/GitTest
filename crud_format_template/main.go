package main

import (
	"gin_test/crud_format_template/config"
	"gin_test/crud_format_template/controller"
	"gin_test/crud_format_template/helper"
	"gin_test/crud_format_template/model"
	"gin_test/crud_format_template/repository"
	"gin_test/crud_format_template/router"
	"gin_test/crud_format_template/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server Started")

	db := config.ConnectDatabase()
	validate := validator.New()
	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepository := repository.NewTagsRepositoryImpl(db)

	tagsService := service.NewTagsRepositoryImpl(tagsRepository, validate)

	tagsController := controller.NewTagsController(tagsService)
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
