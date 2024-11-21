package routes

import (
	"fmt"
	"net/http"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.Users
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not parse request data."})
		return
	}
	user.Id = 1
	err = user.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cloud not create user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Create user"})
}
