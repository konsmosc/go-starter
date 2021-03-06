package services

import (
	"github.com/konsmosc/go-starter/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, User *models.User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, User *models.User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByName(db *gorm.DB, User *models.User, name string) (err error) {
	err = db.Where("username = ?", name).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, User *models.User) (err error) {
	db.Save(User)
	return nil
}
