package repo

import (
	"github.com/Sudhir0302/secure_vault.git/services/auth/config"
	"github.com/Sudhir0302/secure_vault.git/services/auth/models"
)

func Create(user *models.User) (*models.User, error) {
	res := config.DB.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func FindUser(email string) (*models.User, error) {
	user := models.User{}
	res := config.DB.Where("email=?", email).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
