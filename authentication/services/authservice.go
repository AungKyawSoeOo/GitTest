package services

import (
	"gin_test/authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserGet() {

}

func Register(input models.RegisterInput) (int, gin.H) {
	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {

		return http.StatusBadRequest, gin.H{"error": err.Error()}

	}
	return http.StatusOK, gin.H{"message": "registration success"}

}

func Login(input models.LoginInput) (int, gin.H) {

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		return http.StatusBadRequest, gin.H{"error": "username or password is incorrect."}
	}

	return http.StatusOK, gin.H{"token": token}
}
