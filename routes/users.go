package routes

import (
	"example/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse ID"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to signup user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User signed-up"})
}
