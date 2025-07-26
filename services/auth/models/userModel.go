package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	//uuid - 128 bit(32 bit char) unique identifier.collisions are very less
	Id        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"` //char(36) means fixed length to store uuid
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
