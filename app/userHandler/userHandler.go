package userHandler

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/utils"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	_ "github.com/golodash/galidator"
	"github.com/morkid/paginate"
	_ "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var (
	g          = galidator.New()
	customizer = g.Validator(domain.User{})
)

type Handler struct {
	DB *gorm.DB
}

func NewHandlerUser(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) FindALl(c *gin.Context) {

	pg := paginate.New(&paginate.Config{
		DefaultSize: 12,
	})
	var user []domain.User
	model := h.DB.Model(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    pg.With(model).Request(c.Request).Response(&[]domain.User{}),
	})
}

func (h *Handler) FindById(c *gin.Context) {
	var user domain.User
	id := c.Param("id")

	if err := h.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}

func (h *Handler) Create(c *gin.Context) {
	var payload domain.User

	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": customizer.DecryptErrors(err),
		})
		c.Abort()
		return
	}

	rs := h.DB.Where("email=?", payload.Email).First(&payload)
	if rs.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Create new user",
		"data":    newUser,
	})
}

func (h *Handler) Update(c *gin.Context) {
	var user domain.User
	id := c.Param("id")

	if err := h.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found!"})
		return
	}

	payload := user
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": customizer.DecryptErrors(err),
		})
		c.Abort()
		return
	}
	h.DB.Model(&user).Where("id=?", id).Updates(payload)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update data success",
		"data":    payload,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	var user domain.User
	id := c.Param("id")

	if err := h.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found!"})
		return
	}

	payload := user
	h.DB.Delete(&payload, "id=?", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "success delete",
	})
}
