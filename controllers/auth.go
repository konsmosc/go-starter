package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konsmosc/go-starter/database"
	"github.com/konsmosc/go-starter/helpers"
	"github.com/konsmosc/go-starter/models"
	"github.com/konsmosc/go-starter/services"
	"github.com/konsmosc/go-starter/utils"
)

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := services.GetUserByName(database.Db, &user, user.Username)
	if err == nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user.Password = utils.GeneratePassword(user.Password)
	err = services.CreateUser(database.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user_created"})
}

func Login(c *gin.Context) {
	var login Authentication
	var user models.User
	c.BindJSON(&login)
	err := services.GetUserByName(database.Db, &user, login.Username)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	pwdCheck := utils.CheckPassword(login.Password, []byte(user.Password))
	if pwdCheck {
		helpers.Success(c, "user_login", &user)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
