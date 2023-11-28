package userRepository

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/utils"
	"github.com/golodash/galidator"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"time"
)

var (
	g          = galidator.New()
	customizer = g.Validator(domain.User{})
	pg         = paginate.New(&paginate.Config{
		DefaultSize: 12,
	})
)

type UserRepositoryInterface interface {
	Create(payload *domain.User) (*domain.User, error)
}
type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepositoryInterface {
	return &userRepo{DB: db}
}

func (h *userRepo) Create(payload *domain.User) (*domain.User, error) {

	rs := h.DB.Where("email=?", payload.Email).First(&payload)
	if rs.RowsAffected > 0 {
		//Error
		return nil, rs.Error
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		//Error
		return nil, err
	}

	now := time.Now()
	newUser := domain.User{
		Username:      payload.Username,
		Email:         payload.Email,
		Password:      hashedPassword,
		AddressOffice: payload.AddressOffice,
		CreateAt:      now,
		UpdateAt:      now,
	}

	h.DB.Create(&newUser)
	return payload, nil
}
