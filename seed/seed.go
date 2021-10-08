package seed

import (
	"log"

	"github.com/konsmosc/go-starter/models"
	"gorm.io/gorm"
)

var quiz = models.Quiz{
	Category: 1,
	Name:     "CineQuiz",
}

var questions = []models.Question{
	models.Question{
		QuizID:        1,
		Question:      "Who was the last James Bond actor?",
		NumberCorrect: 2,
	},
	models.Question{
		QuizID:        1,
		Question:      "Who is the oldest actor?",
		NumberCorrect: 3,
	},
}

var options = []models.Option{
	models.Option{
		Content:    "Sean Connery",
		QuestionID: 1,
		Correct:    false,
	},
	models.Option{
		Content:    "Daniel Craig",
		QuestionID: 1,
		Correct:    true,
	},
	models.Option{
		Content:    "Pierce Brendan Brosnan",
		QuestionID: 1,
		Correct:    false,
	},
	models.Option{
		Content:    "Timothy Dalton",
		QuestionID: 1,
		Correct:    false,
	},
	models.Option{
		Content:    "Sean Penn",
		QuestionID: 2,
		Correct:    false,
	},
	models.Option{
		Content:    "Val Kilmer",
		QuestionID: 2,
		Correct:    false,
	},
	models.Option{
		Content:    "Hugh Laurie",
		QuestionID: 2,
		Correct:    true,
	},
	models.Option{
		Content:    "Kevin Spacey",
		QuestionID: 2,
		Correct:    false,
	},
}

func Load(db *gorm.DB) {

	if err := db.Debug().Model(&models.Quiz{}).Create(&quiz).Error; err != nil {
		log.Fatalf("cannot seed quizzes table: %v", err)
	}

	for i, _ := range questions {
		err1 := db.Debug().Model(&models.Question{}).Create(&questions[i]).Error
		if err1 != nil {
			log.Fatalf("cannot seed questions table: %v", err1)
		}
	}

	for i, _ := range options {
		err2 := db.Debug().Model(&models.Option{}).Create(&options[i]).Error
		if err2 != nil {
			log.Fatalf("cannot seed options table: %v", err2)
		}
	}
}
