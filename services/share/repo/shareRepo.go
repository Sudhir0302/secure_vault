package repo

import (
	"github.com/Sudhir0302/secure_vault.git/services/share/config"
	"github.com/Sudhir0302/secure_vault.git/services/share/models"
)

func Create(share *models.Share) (*models.Share, error) {
	res := config.DB.Create(&share)
	if res.Error != nil {
		return nil, res.Error
	}
	return share, nil
}

func GetFile(reqbody *models.Share, sharelink string, password string) (*models.Share, error) {
	res := config.DB.Where("share_link=? and password=?", sharelink, password).Find(&reqbody)
	if res.Error != nil {
		return nil, res.Error
	}
	return reqbody, nil
}
