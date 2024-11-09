package di

import (
	"github.com/FabricioCosati/onfly-test/internal/controllers"
	"github.com/FabricioCosati/onfly-test/internal/repository"
	"github.com/FabricioCosati/onfly-test/internal/services"
	"github.com/google/wire"
)

var OrderRepositorySet = wire.NewSet(
	repository.OrderRepositoryInit,
	wire.Bind(new(repository.OrderRepository), new(*repository.OrderRepositoryImpl)),
)

var orderServiceSet = wire.NewSet(
	services.OrderServiceInit,
	wire.Bind(new(services.OrderService), new(*services.OrderServiceImpl)),
)

var orderControllerSet = wire.NewSet(
	controllers.OrderControllerInit,
	wire.Bind(new(controllers.OrderController), new(*controllers.OrderControllerImpl)),
)

var OrderProviders = wire.NewSet(
	OrderRepositorySet,
	orderServiceSet,
	orderControllerSet,
	NewOrderInitialization,
)

type OrderInit struct {
	OrderRepository repository.OrderRepository
	OrderService    services.OrderService
	OrderController controllers.OrderController
}

func NewOrderInitialization(
	OrderRepository repository.OrderRepository,
	orderService services.OrderService,
	orderController controllers.OrderController,
) *OrderInit {
	return &OrderInit{
		OrderRepository: OrderRepository,
		OrderService:    orderService,
		OrderController: orderController,
	}
}
