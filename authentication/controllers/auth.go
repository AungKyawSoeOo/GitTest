package controllers

import (
	"gin_test/authentication/models"
	"gin_test/authentication/services"
	"gin_test/authentication/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// type UserController struct {
// 	AuthService services.Register
// }

func Register(c *gin.Context) {
	var input models.RegisterInput
	// services.Register(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, msg := services.Register(input)
	c.JSON(status, msg)

}
func getErrorMessage(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			switch fieldErr.Field() {
			case "Username":
				return "username is required"
			case "Password":
				return "password is required"
			}
		}
	}
	return err.Error()
}

func Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getErrorMessage(err)})
		return
	}

	status, msg := services.Login(input)
	c.JSON(status, msg)
}

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
