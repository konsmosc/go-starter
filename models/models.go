package models

import (
	_ "fmt"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

type Score struct {
	ID       uint   `json: "id"`
	Username string `json:"username"`
	UserID   uint   `json:"userid"`
	QuizID   uint   `json:"quizid"`
	Value    uint   `json:"value"`
	Attempts uint   `json:"attempts"`
}

type User struct {
	ID         uint    `json:"id"; gorm:"primary_key"`
	Username   string  `json:"username"; gorm:"unique"`
	Password   string  `json:"password"`
	Scores     []Score `json:"scores"`
	TotalScore uint    `json:"totalscore"`
	Admin      bool    `json:"admin"; gorm:"false"`
}

type Option struct {
	ID         uint   `json:"id"; gorm:"primary_key"`
	Content    string `json:"content"`
	QuestionID uint   `json:questionid`
	Correct    bool   `json:"correct"`
}

type Question struct {
	ID            uint     `json:"id"; gorm:"primary_key"`
	QuizID        uint     `json:"quizid"`
	Question      string   `json:"question"`
	Options       []Option `json:"options"`
	NumberCorrect uint     `json:"numcorrect"`
}

type Quiz struct {
	ID        uint       `json:"id"; gorm:"primary_key"`
	Category  uint       `json:"category"`
	Name      string     `json:"name"; gorm:"unique"`
	Questions []Question `json:"questions"`
}
