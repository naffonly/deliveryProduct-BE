package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username      string         `json:"username" gorm:"type:varchar(100);Not null;embedded" binding:"required"`
	Password      string         `json:"password" gorm:"type:varchar(100);Not null;embedded " binding:"required"`
	Email         string         `json:"email" gorm:"type:varchar(100);UNIQUE;Not null;embedded" binding:"required,email"`
	AddressOffice string         `json:"address_office" gorm:"type:varchar(100);Not null;embedded" binding:"required"`
	CreateAt      time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt `json:"delete_at"`
}

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegistrationDTO struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	AddressOffice   string `json:"address_office" binding:"required"`
}
