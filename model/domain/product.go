package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Weight        string         `json:"weight" gorm:"type:varchar(100);Not null" binding:"required"`
	AddressVendor string         `json:"address_vendor" gorm:"type:varchar(100);Not null" binding:"required"`
	NameVendor    string         `json:"name_vendor" gorm:"type:varchar(100);Not null" binding:"required"`
	CreateAt      time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt `json:"delete_at"`
}
