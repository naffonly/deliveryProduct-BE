package logisticRepository

import (
	"deliveryProduct/model/domain"
	"gorm.io/gorm"
	"log"
)

type LogisticRepoInterface interface {
	Create(payload *domain.Logistic) (*domain.Logistic, error)
}

type logisticRepo struct {
	DB *gorm.DB
}

func NewLogisticRepo(db *gorm.DB) LogisticRepoInterface {
	return &logisticRepo{DB: db}
}

func (repository *logisticRepo) Create(payload *domain.Logistic) (*domain.Logistic, error) {

	rs := repository.DB.Create(payload)
	if rs.Error != nil {
		log.Fatal("Create Data Error : ", rs.Error)
		return nil, rs.Error
	}

	return payload, nil
}
