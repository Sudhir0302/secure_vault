package repo

import (
	"github.com/Sudhir0302/secure_vault/services/storage/config"
	"github.com/Sudhir0302/secure_vault/services/storage/models"
)

func Store(store *models.Storage) (*models.Storage, error) {
	res := config.DB.Create(&store)
	if res.Error != nil {
		return nil, res.Error
	}
	return store, nil
}

func GetFile(userid string, fileid string) (*models.Storage, error) {
	data := &models.Storage{}
	res := config.DB.Where("userid=? and id=?", userid, fileid).Find(&data)
	if res.Error != nil {
		return nil, res.Error
	}
	return data, nil
}
