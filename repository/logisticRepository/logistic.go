package logisticRepository

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/model/dto"
	"errors"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"log"
	"strings"
)

var (
	pg = paginate.New(&paginate.Config{
		DefaultSize: 12,
	})
)

type LogisticRepoInterface interface {
	Create(payload *domain.Logistic) (*domain.Logistic, error)
	FindAll(pagination dto.QueryParam) ([]domain.Logistic, error)
	FindById(id string) *domain.Logistic
	Update(payload *domain.Logistic, id string) (*domain.Logistic, error)
	Delete(id string) error
	GetPlatNumber(plat string) (*domain.Logistic, error)
	TotalData() (int64, error)
}
type logisticRepo struct {
	DB *gorm.DB
}

func (repository *logisticRepo) TotalData() (int64, error) {
	var total int64
	var model domain.Logistic
	result := repository.DB.Model(&model).Count(&total)
	if result.Error != nil {
		return -1, result.Error
	}

	return total, nil
}

func (repository *logisticRepo) GetPlatNumber(plat string) (*domain.Logistic, error) {
	var payload domain.Logistic
	rs := repository.DB.Where("plat_number=?", strings.ToUpper(plat)).First(&payload)
	if rs.Error != nil {
		return nil, errors.New("failed FInd Plat Number")
	}
	return &payload, nil
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

func (repository *logisticRepo) FindAll(pagination dto.QueryParam) ([]domain.Logistic, error) {
	var payload []domain.Logistic

	var offset = (pagination.Page - 1) * pagination.Size

	result := repository.DB.Offset(offset).Limit(pagination.Size).Find(&payload)
	if result.Error != nil {
		panic(result.Error)
		return nil, result.Error
	}

	return payload, nil
}

func (repository *logisticRepo) FindById(id string) *domain.Logistic {
	var model *domain.Logistic
	rs := repository.DB.Where("id=?", id).First(&model)
	if rs.Error != nil {
		log.Println("Failed Find Data by id")
		return nil
	}
	return model
}

func (repository *logisticRepo) Update(payload *domain.Logistic, id string) (*domain.Logistic, error) {
	var model domain.Logistic

	rs := repository.DB.Model(model).Where("id=?", id).Updates(payload)
	if rs.Error != nil {
		log.Fatal("Update Data Error : ", rs.Error)
		return nil, rs.Error
	}
	return payload, nil
}

func (repository *logisticRepo) Delete(id string) error {
	var payload domain.Logistic
	if err := repository.DB.Where("id=?", id).First(&payload).Error; err != nil {
		return errors.New("data not found")
	}
	rs := repository.DB.Delete(&payload)
	if rs.Error != nil {
		return errors.New("failed deleted data")
	}
	return nil
}
