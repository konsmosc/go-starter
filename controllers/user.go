package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konsmosc/go-starter/database"
	"github.com/konsmosc/go-starter/models"
	"github.com/konsmosc/go-starter/services"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := database.Db.Create(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var user models.User
	err := services.GetUser(database.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
