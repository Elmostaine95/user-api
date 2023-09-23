package controllers

import (
	"user-api/config"
	"user-api/models"
)

func CreateUser(user *models.User) {
	config.DB.Create(&user)
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return &user, result.Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func UpdateUser(user *models.User) {
	config.DB.Save(&user)
}

func DeleteUser(id uint) {
	config.DB.Delete(&models.User{}, id)
}
