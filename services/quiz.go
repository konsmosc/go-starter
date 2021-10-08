package services

import (
	"github.com/konsmosc/go-starter/models"
	"gorm.io/gorm"
)

func GetQuestions(db *gorm.DB, Quiz *models.Quiz, id string) (Questions []models.Question, err error) {
	var questions []models.Question
	if err := db.Where("id = ?", id).First(Quiz).Error; err != nil {
		return nil, err
	}

	db.Model(Quiz).Association("Questions").Find(&questions)
	for i, question := range questions {
		var options []models.Option
		db.Model(&question).Association("Options").Find(&options)
		questions[i].Options = append(questions[i].Options, options...)
	}

	return questions, nil
}
