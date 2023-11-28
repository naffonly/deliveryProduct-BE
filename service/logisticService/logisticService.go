package logisticService

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/model/dto"
	"deliveryProduct/model/response"
)

type LogisticServiceInterface interface {
	Create(payload *domain.LogisticDto) (*domain.Logistic, error)
	FindAll(pagination dto.QueryParam) ([]domain.Logistic, *response.Pagination, error)
	FindById(id string) (*domain.Logistic, error)
	Update(payload *domain.LogisticDto, id string) (*domain.Logistic, error)
	GetPlatNumber(plat string) (*domain.Logistic, error)
	Delete(id string) error
}
