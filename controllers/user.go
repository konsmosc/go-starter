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
	helpers.Success(c, "user", gin.H{
		"id":       user.ID,
		"username": user.Username,
		"admin":    user.Admin,
	})

}

func UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := c.Params.Get("id")
	err := services.GetUser(database.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = services.UpdateUser(database.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	helpers.Success(c, "user_update", user)
}

func SubmitQuiz(c *gin.Context) {
	var score models.Score
	var score_record models.Score
	var quiz models.Quiz
	var user models.User
	c.BindJSON(&score)
	if err := database.Db.Where("id = ?", score.QuizID).First(&quiz).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		database.Db.Where("id = ?", score.UserID).First(&user)
		if err2 := database.Db.Where("user_id = ? AND quiz_id = ?", score.UserID, score.QuizID).First(&score_record).Error; err2 != nil {
			score.Attempts = 1
			score.Username = user.Username
			database.Db.Create(&score)
			database.Db.Save(&user)
			helpers.Success(c, "submit_quiz", score)
		} else {
			database.Db.Save(&user)
			score_record.Value = score.Value
			score_record.Attempts = score_record.Attempts + 1
			database.Db.Save(&score_record)
			helpers.Success(c, "submit_quiz", score_record)
		}
	}
}

func GetPerformance(c *gin.Context) {
	var scores []models.Score
	var user models.User
	id, _ := c.Params.Get("id")
	if err := database.Db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		database.Db.Model(&user).Association("Scores").Find(&scores)
		user.Scores = append(user.Scores, scores...)
		helpers.Success(c, "get_user_performance", user)
	}
}
