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

type Product struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Weight        int            `json:"Weight" gorm:"type:varchar(100);Not null" binding:"required"`
	AddressVendor string         `json:"address_vendor" gorm:"type:varchar(100);Not null" binding:"required"`
	NameVendor    string         `json:"name_vendor" gorm:"type:varchar(100);Not null" binding:"required"`
	CreateAt      time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt `json:"delete_at"`
}

type Logistic struct {
	ID       uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name     string         `json:"name" gorm:"type:varchar(100);Not null" binding:"required"`
	Address  string         `json:"address" gorm:"type:varchar(100);Not null" binding:"required"`
	CreateAt time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `json:"delete_at"`
}

type TrackingDelivery struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Location      string         `json:"location" gorm:"type:varchar(100);Not null" binding:"required"`
	TransactionID uuid.UUID      `json:"transaction_id" gorm:"type:uuid" binding:"required"`
	CreateAt      time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt `json:"delete_at"`
}

type Transaction struct {
	ID               uuid.UUID          `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID           uuid.UUID          `json:"user_id" gorm:"type:uuid" binding:"required"`
	ProductID        uuid.UUID          `json:"product_id" gorm:"type:uuid" binding:"required"`
	LogisticID       uuid.UUID          `json:"logistic_id" gorm:"type:uuid" binding:"required"`
	Status           string             `json:"status" gorm:"type:varchar(100);Not null" binding:"required"`
	Price            string             `json:"price" gorm:"type:varchar(100);Not null" binding:"required"`
	CreateAt         time.Time          `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt         time.Time          `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt         gorm.DeletedAt     `json:"delete_at"`
	User             User               `gorm:"foreignKey:UserID"`
	Product          Product            `gorm:"foreignKey:ProductID"`
	Logistic         Logistic           `gorm:"foreignKey:LogisticID"`
	TrackingDelivery []TrackingDelivery `gorm:"foreignKey:TransactionID"`
}
