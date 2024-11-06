package controllers

import "github.com/gin-gonic/gin"

func CreateOrderService(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Teste": "Success",
	})
}

func UpdateOrderStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Teste": "Success",
	})
}

func GetOrderService(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Teste": "Success",
	})
}

func GetOrderServices(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Teste": "Success",
	})
}
