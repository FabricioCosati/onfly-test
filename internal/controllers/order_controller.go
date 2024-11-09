package controllers

import (
	"net/http"

	"github.com/FabricioCosati/onfly-test/internal/domain/dto"
	"github.com/FabricioCosati/onfly-test/internal/services"
	cval "github.com/FabricioCosati/onfly-test/internal/validator"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrderService(ctx *gin.Context)
	UpdateOrderStatus(ctx *gin.Context)
	GetOrderService(ctx *gin.Context)
	GetOrderServices(ctx *gin.Context)
}

type OrderControllerImpl struct {
	Service services.OrderService
}

func (impl *OrderControllerImpl) CreateOrderService(ctx *gin.Context) {
	var order dto.OrderRequest

	if err := ctx.ShouldBindBodyWithJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := cval.ValidateRequest(order); err.Errors != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Errors})
		return
	}

	response, err := impl.Service.CreateOrderService(ctx, order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (impl *OrderControllerImpl) UpdateOrderStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Teste": "Success",
	})
}

func (impl *OrderControllerImpl) GetOrderService(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Teste": "Success",
	})
}

func (impl *OrderControllerImpl) GetOrderServices(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Teste": "Success",
	})
}

func OrderControllerInit(service services.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{
		Service: service,
	}
}
