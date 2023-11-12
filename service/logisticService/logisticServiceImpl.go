package logisticService

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/repository/logisticRepository"
	"errors"
	"strings"
)

type logisticServiceImpl struct {
	Repo logisticRepository.LogisticRepoInterface
}

func NewLogisticServiceImpl(repo logisticRepository.LogisticRepoInterface) LogisticServiceInterface {
	return &logisticServiceImpl{Repo: repo}
}

func (l *logisticServiceImpl) Create(payload *domain.LogisticDto) (*domain.Logistic, error) {

	rs := l.Repo.GetPlatNumber(payload.PlatNumber)
	if rs != nil {
		return nil, errors.New("Logistic already exist ")
	}

	var newData = domain.Logistic{
		Name:       payload.Name,
		Address:    payload.Address,
		PlatNumber: strings.ToUpper(payload.PlatNumber),
	}

	result, err := l.Repo.Create(&newData)
	if err != nil {
		return nil, errors.New("insert data menu failed")
	}
	return result, nil
}

func (l *logisticServiceImpl) FindAll(payload []domain.Logistic) (*domain.Logistic, error) {
	//TODO implement me
	panic("implement me")
}

func (l *logisticServiceImpl) FindById(payload *domain.LogisticDto) (*domain.Logistic, error) {
	//TODO implement me
	panic("implement me")
}

func (l *logisticServiceImpl) Update(payload *domain.LogisticDto, id string) (*domain.Logistic, error) {

	rs := l.Repo.GetPlatNumber(payload.PlatNumber)
	if rs != nil {
		return nil, errors.New("Logistic already exist ")
	}

	newPayload := domain.Logistic{
		Name:       payload.Name,
		Address:    payload.Address,
		PlatNumber: strings.ToUpper(payload.PlatNumber),
	}

	rs, err := l.Repo.Update(&newPayload, id)
	if err != nil {
		return nil, errors.New("failed Update Data")
	}
	return rs, nil
}

func (l *logisticServiceImpl) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("id not found")
	}
	err := l.Repo.Delete(id)
	return err
}
