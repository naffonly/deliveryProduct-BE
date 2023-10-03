package transactionHandler

import (
	"deliveryProduct/model/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"time"
)

var (
	g          = galidator.New()
	customizer = g.Validator(domain.Transaction{})
	pg         = paginate.New(&paginate.Config{
		DefaultSize: 12,
	})
)

type Handler struct {
	DB *gorm.DB
}

func NewHandlerUser(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) FindAll(c *gin.Context) {
	var payload []domain.Transaction

	model := h.DB.Preload("User").Preload("Product").Preload("Logistic").Preload("TrackingDelivery").Find(&payload)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    pg.With(model).Request(c.Request).Response(&[]domain.Transaction{}),
	})
}
func (h *Handler) FindById(c *gin.Context) {

}

func (h *Handler) Create(c *gin.Context) {
	var payload domain.TransactionDto
	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": customizer.DecryptErrors(err),
		})
		c.Abort()
		return
	}

	newPayload := domain.Transaction{
		UserID:     payload.UserID,
		LogisticID: payload.LogisticID,
		Product:    payload.Product,
		Status:     payload.Status,
		Price:      payload.Price,
	}
	h.DB.Create(&newPayload)
	c.JSON(http.StatusOK, gin.H{
		"message": "Create new data logistic",
		"data":    newPayload,
	})
}
func (h *Handler) Update(c *gin.Context) {

}
func (h *Handler) Delete(c *gin.Context) {

}

func (h *Handler) UploadImage(c *gin.Context) {
	id := c.Param("id")
	var payload domain.Transaction

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	if err := h.DB.Where("id = ?", id).First(&payload).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found!"})
		return
	}

	ext := filepath.Ext(file.Filename)

	name := fmt.Sprintf("%d%s", time.Now().UnixMilli(), ext)

	if err := c.SaveUploadedFile(file, "./public/storage/"+name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	setFilename := "./public/storage/" + name
	h.DB.Model(&payload).Update("image_delivery", setFilename)

	c.JSON(http.StatusOK, gin.H{
		"message": "upload success",
		"File":    setFilename,
	})

}
