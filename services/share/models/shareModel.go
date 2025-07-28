package models

import (
	"time"

	"github.com/google/uuid"
)

type Share struct {
	ID            uuid.UUID `gorm:"type:char(36);primary key" json:"id"`
	FileId        uuid.UUID `gorm:"type:char(36)" json:"file_id"`
	UserId        uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	ShareLink     string    `gorm:"unique" json:"share_link"`
	Password      string    `json:"password"`
	ExpiryDays    int       `json:"expiry_days"`
	DownloadLimit int       `json:"download_limit"`
	CreatedAt     time.Time `json:"created_at"`
}
