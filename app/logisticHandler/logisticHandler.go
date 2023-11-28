package logisticHandler

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/model/dto"
	"deliveryProduct/model/response"
	"deliveryProduct/service/logisticService"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"net/http"
	"strconv"
)

type logisticHandler struct {
	Service logisticService.LogisticServiceInterface
}

var (
	g          = galidator.New()
	customizer = g.Validator(domain.Logistic{})
)

func NewLogisticHandler(service logisticService.LogisticServiceInterface) LogisticHandlerInterface {
	return &logisticHandler{Service: service}
}

func (h *logisticHandler) FindAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		var queryParam dto.QueryParam

		queryParam.Size, _ = strconv.Atoi(c.Query("size"))
		queryParam.Page, _ = strconv.Atoi(c.Query("page"))
		queryParam.Query = c.Query("query")

		var result []domain.Logistic
		var paginationRes *response.Pagination

		if queryParam.Page < 1 || queryParam.Size < 1 {
			queryParam.Page = 1
			queryParam.Size = 12
		}

		if queryParam.Query != "" {
			rs, err := h.Service.GetPlatNumber(queryParam.Query)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"errors": err.Error(),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg":  "data found",
				"data": rs,
			})
			c.Abort()
			return
		}

		result, paginationRes, err := h.Service.FindAll(queryParam)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}
		data := response.FormatPaginationResponse(
			"Success Find All Data",
			result,
			paginationRes,
		)
		c.JSON(http.StatusOK, data)
	}
}

func (h *logisticHandler) FindById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "id required",
			})
			c.Abort()
			return
		}

		rs, err := h.Service.FindById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success Find Data By Id",
			"data":    rs,
		})
	}
}

func (h *logisticHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var payload domain.LogisticDto
		if err := c.Bind(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		result, err := h.Service.Update(&payload, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messagae": "Successfully Update Data",
			"Data":     result,
		})
	}
}

func (h *logisticHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := h.Service.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Success deleted data",
		})
	}
}

func (h *logisticHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload domain.LogisticDto
		if err := c.Bind(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		result, err := h.Service.Create(&payload)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messagae": "Successfully Create Data",
			"Data":     result,
		})
	}
}
