package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type TrackingDelivery struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Location      string         `json:"location" gorm:"type:varchar(100);Not null" binding:"required"`
	TransactionID uuid.UUID      `json:"transaction_id" gorm:"type:uuid" binding:"required"`
	CreateAt      time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt `json:"delete_at"`
}
