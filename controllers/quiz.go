package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konsmosc/go-starter/database"
	"github.com/konsmosc/go-starter/helpers"
	"github.com/konsmosc/go-starter/models"
	"github.com/konsmosc/go-starter/services"
	"gorm.io/gorm"
)

func GetQuestions(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var quiz models.Quiz
	questions, err := services.GetQuestions(database.Db, &quiz, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	helpers.Success(c, "get_questions", questions)
}
