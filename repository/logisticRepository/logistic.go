package logisticRepository

import (
	"deliveryProduct/model/domain"
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
	FindAll(payload []domain.Logistic) ([]domain.Logistic, error)
	FindById(payload []domain.Logistic) ([]domain.Logistic, error)
	Update(payload *domain.Logistic, id string) (*domain.Logistic, error)
	Delete(id string) error
	GetPlatNumber(plat string) *domain.Logistic
}
type logisticRepo struct {
	DB *gorm.DB
}

func (repository *logisticRepo) GetPlatNumber(plat string) *domain.Logistic {
	var payload domain.Logistic
	rs := repository.DB.Where("plat_number=?", strings.ToUpper(plat)).First(&payload)
	if rs.Error != nil {
		return nil
	}
	return &payload
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

func (repository *logisticRepo) FindAll(payload []domain.Logistic) ([]domain.Logistic, error) {
	//var payload []domain.Logistic
	//rs := repository.DB.Model(&payload)
	//result := pg.With(rs).Request(c.Request).Response(&[]domain.Logistic{})
	//
	//return result, nil
	//TODO implement me
	panic("implement me")
}

func (repository *logisticRepo) FindById(payload []domain.Logistic) ([]domain.Logistic, error) {
	//TODO implement me
	panic("implement me")
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
