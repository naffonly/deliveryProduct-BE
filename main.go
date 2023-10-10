package main

import (
	"deliveryProduct/db"
	"deliveryProduct/middleware"
	"deliveryProduct/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupAppRouter().Run(":8080")
}

func SetupAppRouter() *gin.Engine {
	db := db.InitDB()
	router := gin.Default()

	public := router.Group("/api")
	routes.InitRoutesPublic(db, public)
	protected := router.Group("api/v1")
	protected.Use(middleware.AuthValid)
	routes.InitRoutesProtected(db, protected)

	return router
}
