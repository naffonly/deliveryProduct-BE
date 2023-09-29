package main

import (
	"deliveryProduct/db"
	"deliveryProduct/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupAppRouter().Run(":8080")
}

func SetupAppRouter() *gin.Engine {
	db := db.InitDB()
	router := gin.Default()
	api := router.Group("api/v1")
	routes.InitUserRoutes(db, api)
	return router
}
