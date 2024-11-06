package main

import (
	"fmt"

	"github.com/FabricioCosati/onfly-test/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/order-service", controllers.CreateOrderService)
		api.PATCH("/order-service/:id", controllers.UpdateOrderStatus)
		api.GET("/order-service/:id", controllers.GetOrderService)
		api.GET("/order-services", controllers.GetOrderServices)
	}

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("error on run server: %s", err)
	}
}
