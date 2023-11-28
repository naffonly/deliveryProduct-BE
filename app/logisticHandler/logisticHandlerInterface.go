package logisticHandler

import "github.com/gin-gonic/gin"

type LogisticHandlerInterface interface {
	Create() gin.HandlerFunc
	FindAll() gin.HandlerFunc
	FindById() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
