package db

import (
	"deliveryProduct/model/domain"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func InitDB() *gorm.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Opened to Database")
	Migration(db)
	return db
}

func Migration(db *gorm.DB) {
	fmt.Println("Migration")
	err := db.AutoMigrate(&domain.Logistic{}, &domain.User{}, &domain.Product{}, &domain.Transaction{}, &domain.TrackingDelivery{})
	if err != nil {
		return
	}
	fmt.Print("Seeding")
	seederUser(db)

}

func seederUser(db *gorm.DB) {

	data := domain.User{
		Username:      "admin",
		Password:      "admin",
		Email:         "admin@gmail.com",
		AddressOffice: "Jl. Raya Bogor",
		CreateAt:      time.Now(),
	}
	db.FirstOrCreate(&data)
}
