package controller

import (
	"net/http"

	"github.com/Nathanaelaxl/qresto-api/model"
	"github.com/Nathanaelaxl/qresto-api/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(svc service.OrderService) *OrderController {
	return &OrderController{orderService: svc}
}

func (ctrl *OrderController) CreateOrder(ctx *gin.Context) {
	var input model.Order
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.orderService.PlaceOrder(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Pesanan berhasil dibuat", "data": result})
}
