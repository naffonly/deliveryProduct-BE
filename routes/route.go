package routes

import (
	"deliveryProduct/app/authHandler"
	"deliveryProduct/app/logisticHandler"
	"deliveryProduct/app/trackingHandler"
	"deliveryProduct/app/transactionHandler"
	"deliveryProduct/app/userHandler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutesPublic(db *gorm.DB, router *gin.RouterGroup) {
	authRoutes(db, router)
	pubTrackingRoutes(db, router)
}
func InitRoutesProtected(db *gorm.DB, router *gin.RouterGroup) {
	userRoutes(db, router)
	transactionRoutes(db, router)
	profilRoutes(db, router)
	trackingRoutes(db, router)
}

func InitRoutesLogistic(r *gin.RouterGroup, h logisticHandler.LogisticHandlerInterface) {
	logisticRoutes(r, h)
}

func userRoutes(db *gorm.DB, router *gin.RouterGroup) {
	user := userHandler.NewHandler(db)

	router.GET("/user", user.FindALl)
	router.POST("/user", user.Create)
	router.GET("/user/:id", user.FindById)
	router.PUT("/user/:id", user.Update)
	router.DELETE("/user/:id", user.Delete)

}

func logisticRoutes(router *gin.RouterGroup, handlerInterface logisticHandler.LogisticHandlerInterface) {

	router.GET("/logistic", handlerInterface.FindAll())
	router.GET("/logistic/:id", handlerInterface.FindById())
	router.POST("/logistic", handlerInterface.Create())
	router.PUT("/logistic/:id", handlerInterface.Update())
	router.DELETE("/logistic/:id", handlerInterface.Delete())
}

func transactionRoutes(db *gorm.DB, router *gin.RouterGroup) {
	transaction := transactionHandler.NewHandler(db)
	router.GET("/transaction", transaction.FindAll)
	router.GET("/transaction/:id", transaction.FindById)
	router.POST("/transaction", transaction.Create)
	router.PUT("/transaction/upload-image/:id", transaction.UploadImage)
	router.GET("/transaction/image/:id", transaction.GetFile)
	router.PUT("/transaction/:id", transaction.Update)
	router.DELETE("/transaction/:id", transaction.Delete)
}

func trackingRoutes(db *gorm.DB, router *gin.RouterGroup) {
	tracking := trackingHandler.NewHandler(db)
	router.POST("/tracking", tracking.Create)
	router.PUT("/tracking/:id", tracking.Update)
	router.DELETE("/tracking/:id", tracking.Delete)
}
func profilRoutes(db *gorm.DB, router *gin.RouterGroup) {
	auth := authHandler.NewHandler(db)
	router.GET("/profil", auth.CurrentUser)
}

func authRoutes(db *gorm.DB, router *gin.RouterGroup) {
	auth := authHandler.NewHandler(db)
	transaction := transactionHandler.NewHandler(db)

	router.GET("/transaction/image/:id", transaction.GetFile)
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
}

func pubTrackingRoutes(db *gorm.DB, router *gin.RouterGroup) {
	transaction := transactionHandler.NewHandler(db)

	router.GET("/check/:code", transaction.GetTransactionByCode)
}
