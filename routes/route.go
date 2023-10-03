package routes

import (
	"deliveryProduct/app/logisticHandler"
	"deliveryProduct/app/userHandler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB, router *gin.RouterGroup) {
	userRoutes(db, router)
	logisticRoutes(db, router)
}

func userRoutes(db *gorm.DB, router *gin.RouterGroup) {
	user := userHandler.NewHandlerUser(db)

	router.GET("/user", user.FindALl)
	router.POST("/user", user.Create)
	router.GET("/user/:id", user.FindById)
	router.PUT("/user/:id", user.Update)
	router.DELETE("/user/:id", user.Delete)

}

func logisticRoutes(db *gorm.DB, router *gin.RouterGroup) {
	logistic := logisticHandler.NewHandlerUser(db)

	router.GET("/logistic", logistic.FindAll)
	router.GET("/logistic/:id", logistic.FindById)
	router.POST("/logistic", logistic.Create)
	router.PUT("/logistic/:id", logistic.Update)
	router.DELETE("/logistic/:id", logistic.Delete)
}
