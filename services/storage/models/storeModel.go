package models

import (
	"time"

	"github.com/google/uuid"
)

type Storage struct {
	//schema
	ID            uuid.UUID `gorm:"type:char(36);primary key" json:"id"`
	Userid        uuid.UUID `gorm:"type:char(36)" json:"userid"`
	FileName      string    `json:"file_name"`
	FileSize      int64     `json:"file_size"`
	Mime_type     string    `json:"mime_type"`
	EncryptedData []byte    `json:"encrypted_data"`
	UploadedAt    time.Time `json:"uploaded_at"`
}
