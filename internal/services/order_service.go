package services

import (
	"errors"
	"strconv"

	"github.com/FabricioCosati/onfly-test/internal/domain/dao"
	"github.com/FabricioCosati/onfly-test/internal/domain/dto"
	ce "github.com/FabricioCosati/onfly-test/internal/errors"
	"github.com/FabricioCosati/onfly-test/internal/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrderService(ctx *gin.Context, order dto.OrderRequestPost) (dao.Order, ce.CustomError)
	UpdateOrderStatus(ctx *gin.Context, order dto.OrderRequestPatch) (dao.Order, ce.CustomError)
	GetOrderById(ctx *gin.Context) (dao.Order, ce.CustomError)
	GetOrders(ctx *gin.Context) (dao.OrderCollection, ce.CustomError)
}

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
}

func (impl *OrderServiceImpl) CreateOrderService(ctx *gin.Context, orderRequest dto.OrderRequestPost) (dao.Order, ce.CustomError) {
	order := dao.Order{
		RequesterName: orderRequest.RequesterName,
		Destination:   orderRequest.Destination,
		GoingDate:     orderRequest.GoingDate.Time,
		ReturnDate:    orderRequest.ReturnDate.Time,
		Status:        orderRequest.Status,
	}

	order, err := impl.OrderRepository.CreateOrder(&order)

	if err != nil {
		return order, *ce.InternalServerError()
	}

	return order, ce.CustomError{}
}

func (impl *OrderServiceImpl) UpdateOrderStatus(ctx *gin.Context, orderRequest dto.OrderRequestPatch) (dao.Order, ce.CustomError) {
	orderId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return dao.Order{}, ce.CustomError{}
	}

	order := dao.Order{
		Status: orderRequest.Status,
	}

	order, err = impl.OrderRepository.UpdateStatusOrder(&order, orderId)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return order, *ce.NotFound(err)
		default:
			return order, *ce.InternalServerError()
		}
	}

	return order, ce.CustomError{}
}

func (impl *OrderServiceImpl) GetOrderById(ctx *gin.Context) (dao.Order, ce.CustomError) {
	orderId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return dao.Order{}, ce.CustomError{}
	}

	order, err := impl.OrderRepository.GetOrderById(orderId)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return order, *ce.NotFound(err)
		default:
			return order, *ce.InternalServerError()
		}
	}

	return order, ce.CustomError{}
}

func (impl *OrderServiceImpl) GetOrders(ctx *gin.Context) (dao.OrderCollection, ce.CustomError) {
	orders, err := impl.OrderRepository.GetOrders()

	if err != nil {
		return orders, *ce.InternalServerError()
	}

	return orders, ce.CustomError{}
}

func OrderServiceInit(repository repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{
		OrderRepository: repository,
	}
}
