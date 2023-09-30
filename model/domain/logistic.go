package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Logistic struct {
	ID         uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name       string         `json:"name" gorm:"type:varchar(100);Not null" binding:"required"`
	Address    string         `json:"address" gorm:"type:varchar(100);Not null" binding:"required"`
	PlatNumber string         `json:"plat_number" gorm:"type:varchar(20);Not Null" binding:"required"`
	CreateAt   time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt   time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt   gorm.DeletedAt `json:"delete_at"`
}

type LogisticDto struct {
	ID         uuid.UUID      `json:"id"`
	Name       string         `json:"name"`
	Address    string         `json:"address"`
	PlatNumber string         `json:"plat_number"`
	CreateAt   time.Time      `json:"create_at"`
	UpdateAt   time.Time      `json:"update_at"`
	DeleteAt   gorm.DeletedAt `json:"delete_at"`
}
