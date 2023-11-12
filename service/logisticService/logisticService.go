package logisticService

import "deliveryProduct/model/domain"

type LogisticServiceInterface interface {
	Create(payload *domain.LogisticDto) (*domain.Logistic, error)
	FindAll(payload []domain.Logistic) (*domain.Logistic, error)
	FindById(payload *domain.LogisticDto) (*domain.Logistic, error)
	Update(payload *domain.LogisticDto, id string) (*domain.Logistic, error)
	Delete(id string) error
}
