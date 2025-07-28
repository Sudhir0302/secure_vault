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

func GetFile(reqbody *models.Share, sharelink string, password string) (int64, error) {
	res := config.DB.Where("share_link=? and password=?", sharelink, password).Find(&reqbody)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DecreLimit(share *models.Share) error {
	// .Model(&models.Share{}) tells GORM what table to use.
	res := config.DB.Model(&models.Share{}).Where("id=?", share.ID).Update("download_limit", share.DownloadLimit)
	return res.Error
}

func Delete(share *models.Share) error {
	res := config.DB.Model(&models.Share{}).Delete("id=?", share.ID)
	return res.Error
}
