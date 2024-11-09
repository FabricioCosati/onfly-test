package routes

import (
	"github.com/FabricioCosati/onfly-test/internal/di"
)

func InitOrderRoutes(app *di.AppInit) {
	controllers := app.OrderInit.OrderController

	api := app.Server.Engine.Group("/api")
	{
		api.POST("/order-service", controllers.CreateOrderService)
		api.PATCH("/order-service/:id", controllers.UpdateOrderStatus)
		api.GET("/order-service/:id", controllers.GetOrderService)
		api.GET("/order-services", controllers.GetOrderServices)
	}
}
