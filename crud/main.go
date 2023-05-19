package main

import (
	"gin_test/crud/controllers"
	"gin_test/crud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.LoadHTMLGlob("templates/**/**")
	r.Static("/static", "./static")
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
