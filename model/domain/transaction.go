package domain

import (
	_ "github.com/joho/godotenv"
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID               uuid.UUID          `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID           uuid.UUID          `json:"user_id" gorm:"type:uuid" binding:"required"`
	ProductID        uuid.UUID          `json:"product_id" gorm:"type:uuid"`
	LogisticID       uuid.UUID          `json:"logistic_id" gorm:"type:uuid" binding:"required"`
	Status           string             `json:"status" gorm:"type:varchar(100);Not null" binding:"required"`
	Price            string             `json:"price" gorm:"type:varchar(100);Not null" binding:"required"`
	ImageDelivery    string             `json:"image_delivery" gorm:"type:varchar(100);"`
	CreateAt         time.Time          `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt         time.Time          `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt         gorm.DeletedAt     `json:"delete_at"`
	User             User               `gorm:"foreignKey:UserID"`
	Product          Product            `gorm:"foreignKey:ProductID"`
	Logistic         Logistic           `gorm:"foreignKey:LogisticID"`
	TrackingDelivery []TrackingDelivery `gorm:"foreignKey:TransactionID"`
}

type TransactionDto struct {
	UserID     uuid.UUID `json:"user_id" binding:"required"`
	LogisticID uuid.UUID `json:"logistic_id" binding:"required"`
	Status     string    `json:"status" binding:"required"`
	Price      string    `json:"price" binding:"required"`
	Product    Product   `json:"product"`
}
