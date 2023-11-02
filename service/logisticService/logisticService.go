package logisticService

import "deliveryProduct/model/domain"

type LogisticServiceInterface interface {
	Create(payload domain.LogisticDto) (*domain.Logistic, error)
	FindAll(payload []domain.Logistic) (*domain.Logistic, error)
	FindById(payload domain.LogisticDto) (*domain.Logistic, error)
	Update(payload domain.LogisticDto) (*domain.Logistic, error)
	Delete(payload domain.LogisticDto) (*domain.Logistic, error)
}
