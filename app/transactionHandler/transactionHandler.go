package transactionHandler

import (
	"deliveryProduct/model/domain"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
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

}
func (h *Handler) FindById(c *gin.Context) {

}

func (h *Handler) Create(c *gin.Context) {

}
func (h *Handler) Update(c *gin.Context) {

}
func (h *Handler) Delete(c *gin.Context) {

}
