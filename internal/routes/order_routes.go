package routes

import (
	"github.com/FabricioCosati/onfly-test/internal/di"
	"github.com/FabricioCosati/onfly-test/internal/middlewares"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func InitOrderRoutes(app *di.AppInit) {
	controllers := app.OrderInit.OrderController

	api := app.Server.Engine.Group("/api")
	{
		api.Use(middlewares.OperationLogs("operation"))
		api.Use(otelgin.Middleware("onfly-orders"))
		api.Use(middlewares.OtelMiddleware())

		api.POST("/order-service", controllers.CreateOrderService)
		api.PATCH("/order-service/:id", controllers.UpdateOrderStatus)
		api.GET("/order-service/:id", controllers.GetOrderById)
		api.GET("/order-services", controllers.GetOrders)
	}
}
