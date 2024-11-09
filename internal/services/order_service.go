package services

import (
	"github.com/FabricioCosati/onfly-test/internal/domain/dao"
	"github.com/FabricioCosati/onfly-test/internal/domain/dto"
	"github.com/FabricioCosati/onfly-test/internal/repository"
	"github.com/gin-gonic/gin"
)

type OrderService interface {
	CreateOrderService(ctx *gin.Context, order dto.OrderRequest) (dao.Order, error)
}

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
}

func (impl *OrderServiceImpl) CreateOrderService(ctx *gin.Context, orderRequest dto.OrderRequest) (dao.Order, error) {
	order := dao.Order{
		RequesterName: orderRequest.RequesterName,
		Destination:   orderRequest.Destination,
		GoingDate:     orderRequest.GoingDate.Time,
		ReturnDate:    orderRequest.ReturnDate.Time,
		Status:        orderRequest.Status,
	}

	return impl.OrderRepository.CreateOrder(&order)
}

func OrderServiceInit(repository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		OrderRepository: repository,
	}
}
