package routes

import (
	"deliveryProduct/app/authHandler"
	"deliveryProduct/app/logisticHandler"
	"deliveryProduct/app/transactionHandler"
	"deliveryProduct/app/userHandler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutesPublic(db *gorm.DB, router *gin.RouterGroup) {
	authRoutes(db, router)

}
func InitRoutesProtected(db *gorm.DB, router *gin.RouterGroup) {
	userRoutes(db, router)
	logisticRoutes(db, router)
	transactionRoutes(db, router)
	profilRoutes(db, router)
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

func transactionRoutes(db *gorm.DB, router *gin.RouterGroup) {
	transaction := transactionHandler.NewHandlerUser(db)
	router.GET("/transaction", transaction.FindAll)
	router.GET("/transaction/:id", transaction.FindById)
	router.POST("/transaction", transaction.Create)
	router.PUT("/transaction/upload-image/:id", transaction.UploadImage)
	router.GET("/transaction/image/:id", transaction.GetFile)
	router.PUT("/transaction/:id", transaction.Update)
	router.DELETE("/transaction/:id", transaction.Delete)
}
func profilRoutes(db *gorm.DB, router *gin.RouterGroup) {
	auth := authHandler.NewHandlerUser(db)
	router.GET("/profil", auth.CurrentUser)
}

func authRoutes(db *gorm.DB, router *gin.RouterGroup) {
	auth := authHandler.NewHandlerUser(db)
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
}
