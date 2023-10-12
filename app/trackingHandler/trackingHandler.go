package trackingHandler

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
}

func NewHandlerUser(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) Create() {

}
