package db

import (
	"deliveryProduct/config"
	"deliveryProduct/model/domain"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitDB(config config.Config) *gorm.DB {

	psql := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBPort)

	db, err := gorm.Open(postgres.Open(psql), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Opened to Database")
	Migration(db)
	return db
}

func Migration(db *gorm.DB) {

	//fmt.Println("Migration")
	//err := db.AutoMigrate(&domain.Logistic{}, &domain.User{}, &domain.Product{}, &domain.Transaction{}, &domain.TrackingDelivery{})
	//fmt.Print("Drop Table")
	//err := db.Migrator().DropTable(&domain.TrackingDelivery{}, &domain.Transaction{}, &domain.Product{})
	//if err != nil {
	//	return
	//}
	//fmt.Print("Seeding")

	fmt.Println("Migration")

	errs := db.AutoMigrate(&domain.Logistic{}, &domain.User{}, &domain.Product{}, &domain.Transaction{}, &domain.TrackingDelivery{})
	if errs != nil {
		return
	}
	//fmt.Print("Seeding")
	//seederUser(db)

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
