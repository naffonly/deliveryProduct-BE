package logisticHandler

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/service/logisticService"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/morkid/paginate"
	"net/http"
)

var (
	g          = galidator.New()
	customizer = g.Validator(domain.Logistic{})
	pg         = paginate.New(&paginate.Config{
		DefaultSize: 12,
	})
)

type logisticHandler struct {
	Service logisticService.LogisticServiceInterface
}

func NewLogisticHandler(service logisticService.LogisticServiceInterface) LogisticHandlerInterface {
	return &logisticHandler{
		Service: service,
	}
}

func (h *logisticHandler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var result []domain.Logistic
		data, err := h.Service.FindAll(result)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    data})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    data})
	}
}

func (h *logisticHandler) FindById() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *logisticHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *logisticHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func (h *logisticHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//func (h *Handler) FindAll(c *gin.Context) {
//	var payload []domain.Logistic
//	model := h.DB.Model(&payload)
//	c.JSON(http.StatusOK, gin.H{
//		"message": "success",
//		"data":    pg.With(model).Request(c.Request).Response(&[]domain.Logistic{}),
//	})
//}

//func (h *Handler) FindById(c *gin.Context) {
//	id := c.Param("id")
//	var payload domain.Logistic
//
//	if err := h.DB.Where("id = ?", id).First(&payload).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "user not found!"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "success",
//		"data":    payload,
//	})
//}

//func (h *Handler) Create(c *gin.Context) {
//	var payload domain.Logistic
//	err := c.BindJSON(&payload)
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": customizer.DecryptErrors(err),
//		})
//		c.Abort()
//		return
//	}
//
//	rs := h.DB.Where("plat_number=?", strings.ToUpper(payload.PlatNumber)).First(&payload)
//	if rs.RowsAffected > 0 {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": "Logistic already exist",
//		})
//		return
//	}
//	now := time.Now()
//
//	newPayload := domain.Logistic{
//		Name:       payload.Name,
//		Address:    payload.Address,
//		PlatNumber: strings.ToUpper(payload.PlatNumber),
//		CreateAt:   now,
//		UpdateAt:   now,
//	}
//
//	h.DB.Create(&newPayload)
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Create new data logistic",
//		"data":    newPayload,
//	})
//
//}
//func (h *Handler) Update(c *gin.Context) {
//	id := c.Param("id")
//	var data domain.Logistic
//
//	if err := h.DB.Where("id = ?", id).First(&data).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "user not found!"})
//		return
//	}
//
//	payload := data
//	err := c.BindJSON(&payload)
//
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"error": customizer.DecryptErrors(err),
//		})
//		c.Abort()
//		return
//	}
//
//	newPayload := domain.Logistic{
//		Name:       payload.Name,
//		Address:    payload.Address,
//		PlatNumber: strings.ToUpper(payload.PlatNumber),
//	}
//	h.DB.Model(&payload).Where("id=?", id).Updates(newPayload)
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Update data success",
//		"data":    payload,
//	})
//
//}
//func (h *Handler) Delete(c *gin.Context) {
//	id := c.Param("id")
//	payload := domain.Logistic{}
//
//	if err := h.DB.Where("id=?", id).First(&payload).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{
//			"error": "Data not found",
//		})
//		return
//	}
//
//	h.DB.Delete(&payload)
//	c.JSON(http.StatusOK, gin.H{
//		"message": "deleted data success",
//	})
//}
