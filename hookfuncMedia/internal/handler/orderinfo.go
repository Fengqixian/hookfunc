package handler

import (
	"github.com/gin-gonic/gin"
	"hookfunc-media/internal/service"
)

type OrderInfoHandler struct {
	*Handler
	orderInfoService service.OrderInfoService
}

func NewOrderInfoHandler(handler *Handler, orderInfoService service.OrderInfoService) *OrderInfoHandler {
	return &OrderInfoHandler{
		Handler:          handler,
		orderInfoService: orderInfoService,
	}
}

func (h *OrderInfoHandler) GetOrderInfo(ctx *gin.Context) {

}
