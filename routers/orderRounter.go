package routers

import (
	"Assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.PUT("/orders/:orderId", controllers.UpdateOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return router

}
