package database

import (
	"users-books-api-testing/config"
	"users-books-api-testing/middlewares"
	"users-books-api-testing/models"
)

func CreateUser(user *models.Users) error {
	if err := config.DB.Table("users").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() (interface{}, error) {
	var users []models.Users

	if err := config.DB.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var user models.Users

	if err := config.DB.Table("users").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserById(id int, user *models.Users) error {
	var users models.Users
	if err := config.DB.Table("users").First(&users, id).Error; err != nil {
		return err
	}
	err := config.DB.Table("users").Where("id = ?", id).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id int) error {
	var user models.Users
	if err := config.DB.Table("users").Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.Users) (interface{}, error){
	var err error
	err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).Error
	if err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
