package trackingHandler

import (
	"deliveryProduct/model/domain"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"net/http"
)

var (
	g          = galidator.New()
	customizer = g.Validator(domain.TrackingDelivery{})
	pg         = paginate.New(&paginate.Config{
		DefaultSize: 12,
	})
)

type Handler struct {
	DB *gorm.DB
}

func NewHandlerTracking(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) Create(c *gin.Context) {
	var payload domain.TrackingDelivery
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": customizer.DecryptErrors(err),
		})
		c.Abort()
		return
	}

	newPayload := domain.TrackingDelivery{
		Location:      payload.Location,
		TransactionID: payload.TransactionID,
	}

	h.DB.Create(&newPayload)
	c.JSON(http.StatusOK, gin.H{
		"message": "Create new data logistic",
		"data":    newPayload,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	payload := domain.TrackingDelivery{}

	if err := h.DB.Where("id = ?", id).First(&payload).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found!"})
		return
	}

	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": customizer.DecryptErrors(err),
		})
		c.Abort()
		return
	}

	newPayload := domain.TrackingDelivery{
		Location:      payload.Location,
		TransactionID: payload.TransactionID,
	}

	h.DB.Model(&payload).Where("id=?", id).Updates(newPayload)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update data success",
		"data":    payload,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	payload := domain.TrackingDelivery{}

	if err := h.DB.Where("id=?", id).First(&payload).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found",
		})
		return
	}

	h.DB.Delete(&payload)
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted data success",
	})
}
