package main

import (
	configHandler "deliveryProduct/config"
	"deliveryProduct/middleware"
	"deliveryProduct/routes"
	dbHandler "deliveryProduct/utils/db"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupAppRouter().Run(":8080")
}

func SetupAppRouter() *gin.Engine {
	config := configHandler.InitConfig()
	db := dbHandler.InitDB(*config)
	router := gin.Default()

	public := router.Group("/api")
	routes.InitRoutesPublic(db, public)
	protected := router.Group("api/v1")
	protected.Use(middleware.AuthValid)
	routes.InitRoutesProtected(db, protected)

	return router
}
