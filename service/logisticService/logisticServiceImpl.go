package logisticService

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/model/dto"
	"deliveryProduct/model/response"
	"deliveryProduct/repository/logisticRepository"
	"errors"
	"fmt"
	"strings"
)

type logisticServiceImpl struct {
	Repo logisticRepository.LogisticRepoInterface
}

func NewLogisticServiceImpl(repo logisticRepository.LogisticRepoInterface) LogisticServiceInterface {
	return &logisticServiceImpl{Repo: repo}
}

func (l *logisticServiceImpl) Create(payload *domain.LogisticDto) (*domain.Logistic, error) {

	rs, _ := l.Repo.GetPlatNumber(payload.PlatNumber)
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

func (l *logisticServiceImpl) FindAll(pagination dto.QueryParam) ([]domain.Logistic, *response.Pagination, error) {
	rs, err := l.Repo.FindAll(pagination)
	if err != nil {
		return nil, nil, errors.New("get data logistic failed")
	}
	fmt.Println()
	var logisticRes []domain.Logistic

	for _, value := range rs {
		logisticRes = append(logisticRes, value)
	}
	total, err := l.Repo.TotalData()
	if err != nil {
		return nil, nil, errors.New("get total menu failed")
	}

	var DataResponse = &response.Pagination{
		Page:       pagination.Page,
		PageSize:   pagination.Size,
		TotalItems: total,
	}

	return logisticRes, DataResponse, nil
}

func (l *logisticServiceImpl) FindById(id string) (*domain.Logistic, error) {
	rs := l.Repo.FindById(id)
	if rs == nil {
		return nil, errors.New("data Logistic not found")
	}
	return rs, nil
}

func (l *logisticServiceImpl) Update(payload *domain.LogisticDto, id string) (*domain.Logistic, error) {

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

func (l *logisticServiceImpl) GetPlatNumber(plat string) (*domain.Logistic, error) {

	rs, err := l.Repo.GetPlatNumber(plat)
	if err != nil {
		return nil, err
	}
	return rs, nil
}
