package repo

import (
	"github.com/Sudhir0302/secure_vault.git/services/storage/config"
	"github.com/Sudhir0302/secure_vault.git/services/storage/models"
)

func Store(store *models.Storage) (*models.Storage, error) {
	res := config.DB.Create(&store)
	if res.Error != nil {
		return nil, res.Error
	}
	return store, nil
}

func GetFile(userid string, filename string) (*models.Storage, error) {
	data := &models.Storage{}
	res := config.DB.Where("userid=? and file_name=?", userid, filename).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return data, nil
}
