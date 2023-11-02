package logisticService

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/repository/logisticRepository"
)

type logisticServiceImpl struct {
	Repo logisticRepository.LogisticRepoInterface
}

func NewLogisticServiceImpl(repo logisticRepository.LogisticRepoInterface) LogisticServiceInterface {
	return &logisticServiceImpl{Repo: repo}
}

func (service *logisticServiceImpl) FindAll(payload domain.LogisticDto) (*domain.Logistic, error) {
	//TODO implement me
	panic("implement me")
}

func (service *logisticServiceImpl) FindById(payload domain.LogisticDto) (*domain.Logistic, error) {
	//TODO implement me
	panic("implement me")
}

func (service *logisticServiceImpl) Update(payload domain.LogisticDto) (*domain.Logistic, error) {
	//TODO implement me
	panic("implement me")
}

func (service *logisticServiceImpl) Delete(payload domain.LogisticDto) (*domain.Logistic, error) {
	//TODO implement me
	panic("implement me")
}

func (service *logisticServiceImpl) Create(payload domain.LogisticDto) (*domain.Logistic, error) {
	return nil, nil
}
