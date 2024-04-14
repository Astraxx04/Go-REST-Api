package routes

import (
	"net/http"
	"rest-api/models"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messaage": "Could not create user!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.ValidateCredentials()

	if(err != nil) {
		context.JSON(http.StatusUnauthorized, gin.H{"messaage": "Could not authenticate user!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if(err != nil) {
		context.JSON(http.StatusInternalServerError, gin.H{"messaage": "Could not authenticate user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}