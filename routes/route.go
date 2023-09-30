package routes

import (
	"deliveryProduct/app/userHandler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, router *gin.RouterGroup) {
	user := userHandler.NewHandlerUser(db)

	router.GET("/user", user.FindALl)
	router.POST("/user", user.Create)
	router.GET("/user/:id", user.FindById)
	router.PUT("/user/:id", user.Update)
	router.DELETE("/user/:id", user.Delete)

}

func InitLogisticRoutes(db *gorm.DB, router *gin.RouterGroup) {

}
