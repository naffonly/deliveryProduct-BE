package main

import (
	"deliveryProduct/app/logisticHandler"
	configHandler "deliveryProduct/config"
	"deliveryProduct/middleware"
	"deliveryProduct/repository/logisticRepository"
	"deliveryProduct/routes"
	logisticService2 "deliveryProduct/service/logisticService"
	dbHandler "deliveryProduct/utils/db"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupAppRouter().Run(":8080")
}

func SetupAppRouter() *gin.Engine {
	config := configHandler.InitConfig()
	router := gin.Default()

	db := dbHandler.InitDB(*config)

	logisticRepo := logisticRepository.NewLogisticRepo(db)
	logisticService := logisticService2.NewLogisticServiceImpl(logisticRepo)
	logisticHadler := logisticHandler.NewLogisticHandler(logisticService)

	public := router.Group("/api")
	routes.InitRoutesPublic(db, public)
	//Testing Clean Architecture
	routes.InitRoutesLogistic(public, logisticHadler)

	protected := router.Group("api/v1")
	protected.Use(middleware.AuthValid)
	routes.InitRoutesProtected(db, protected)

	return router
}
